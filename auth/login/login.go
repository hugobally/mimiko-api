package login

import (
	"context"
	"errors"
	"github.com/hugobally/mimiko/backend/prisma"
	"github.com/hugobally/mimiko/backend/spotify"
	"io/ioutil"
	"net/http"
	"time"
)

func (h *Handler) LoginSampleSession(w http.ResponseWriter, r *http.Request) {
	if h.Config.Env == "DEV" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	}

	uuid, err := h.ParseRequest(w, r)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	token, err := h.Spotify.CreateClientCredentialsToken()
	if err != nil {
		h.UnprocessableResponse(w, err)
		return
	}

	user, err := h.UpsertSampleSessionUser(uuid, token, r.Context())
	if err != nil {
		h.UnprocessableResponse(w, err)
		return
	}

	err = h.SetLoginCookie(w, user)
	if err != nil {
		h.Logger.Println("error on jwt creation", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.Logger.Println("successful login for user", user.ID)
}

// Uncomment to re-enable spotify login -- also uncomment the refreshToken piece of code in mutation/token.go
//
//	func (h *Handler) LoginAuthCode(w http.ResponseWriter, r *http.Request) {
//		if h.Config.Env == "DEV" {
//			w.Header().Set("Access-Control-Allow-Credentials", "true")
//			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
//		}
//
//		if r.Method == http.MethodGet {
//			h.SpotifyRedirect(w, r)
//			return
//		}
//
//		authCode, err := h.ParseRequest(w, r)
//		if err != nil {
//			_, _ = w.Write([]byte(err.Error()))
//			return
//		}
//
//		token, err := h.Spotify.CreateAuthCodeToken(*authCode)
//		if err != nil {
//			h.UnprocessableResponse(w, err)
//			return
//		}
//
//		spotifyUser, err := h.Spotify.GetUser(token.AccessToken)
//		if err != nil {
//			h.UnprocessableResponse(w, err)
//			return
//		}
//
//		user, err := h.UpsertSpotifyUser(spotifyUser, token, r.Context())
//		if err != nil {
//			h.UnprocessableResponse(w, err)
//			return
//		}
//
//		err = h.SetLoginCookie(w, user)
//		if err != nil {
//			h.Logger.Println("error on jwt creation", err)
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//
//		h.Logger.Println("successful login for user", user.ID)
//	}
func (h *Handler) ParseRequest(w http.ResponseWriter, r *http.Request) (*string, error) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil, h.LogError("invalid method", nil)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, h.LogError("error while reading body", err)
	}
	bodyStr := string(body)
	if bodyStr == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, h.LogError("auth code not provided", nil)
	}

	return &bodyStr, nil
}

// Could not use actual UPSERT statement because (remote) UserId is not unique
func (h *Handler) UpsertSpotifyUser(u *spotify.UserResponse, d *spotify.TokenResponse, ctx context.Context) (*prisma.User, error) {
	appType := prisma.AppTypeSpotify

	apps, err := h.Prisma.LinkedApps(&prisma.LinkedAppsParams{
		Where: &prisma.LinkedAppWhereInput{
			Type:   &appType,
			UserId: &u.Id,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, h.LogError("internal server error", err)
	}

	exp := time.Now().Add(time.Duration(d.ExpiresIn) * time.Second)
	expStr := exp.Format(time.RFC3339)

	newUser := &prisma.User{}

	if len(apps) == 0 {
		newUser, err = h.Prisma.CreateLinkedApp(prisma.LinkedAppCreateInput{
			UserId:   u.Id,
			Username: &u.DisplayName,
			Type:     prisma.AppTypeSpotify,
			User: prisma.UserCreateOneWithoutLinkedAppsInput{
				Create: &prisma.UserCreateWithoutLinkedAppsInput{
					Username: &u.DisplayName,
				},
			},
			AccessToken:  &d.AccessToken,
			TokenExpiry:  &expStr,
			RefreshToken: &d.RefreshToken,
		}).User().Exec(ctx)
	} else {
		app := apps[0]
		newUser, err = h.Prisma.UpdateLinkedApp(prisma.LinkedAppUpdateParams{
			Data: prisma.LinkedAppUpdateInput{
				AccessToken:  &d.AccessToken,
				TokenExpiry:  &expStr,
				RefreshToken: &d.RefreshToken,
			},
			Where: prisma.LinkedAppWhereUniqueInput{
				ID: &app.ID,
			},
		}).User().Exec(ctx)
	}

	if err != nil {
		return nil, h.LogError("internal server error", err)
	}

	return newUser, nil
}

func (h *Handler) UpsertSampleSessionUser(uuid *string, d *spotify.TokenResponse, ctx context.Context) (*prisma.User, error) {
	appType := prisma.AppTypeSpotify

	apps, err := h.Prisma.LinkedApps(&prisma.LinkedAppsParams{
		Where: &prisma.LinkedAppWhereInput{
			Type:     &appType,
			Username: uuid,
		},
	}).Exec(ctx)
	if err != nil {
		return nil, h.LogError("internal server error", err)
	}

	exp := time.Now().Add(time.Duration(d.ExpiresIn) * time.Second)
	expStr := exp.Format(time.RFC3339)

	newUser := &prisma.User{}

	if len(apps) == 0 {
		newUser, err = h.Prisma.CreateLinkedApp(prisma.LinkedAppCreateInput{
			Username: uuid,
			Type:     prisma.AppTypeSpotify,
			User: prisma.UserCreateOneWithoutLinkedAppsInput{
				Create: &prisma.UserCreateWithoutLinkedAppsInput{
					Username: uuid,
				},
			},
			AccessToken:  &d.AccessToken,
			TokenExpiry:  &expStr,
			RefreshToken: nil,
		}).User().Exec(ctx)
	} else {
		app := apps[0]
		newUser, err = h.Prisma.UpdateLinkedApp(prisma.LinkedAppUpdateParams{
			Data: prisma.LinkedAppUpdateInput{
				AccessToken:  &d.AccessToken,
				TokenExpiry:  &expStr,
				RefreshToken: nil,
			},
			Where: prisma.LinkedAppWhereUniqueInput{
				ID: &app.ID,
			},
		}).User().Exec(ctx)
	}

	if err != nil {
		return nil, h.LogError("internal server error", err)
	}

	return newUser, nil
}

func (h *Handler) SetLoginCookie(w http.ResponseWriter, user *prisma.User) error {
	exp := time.Now().Add(8750 * time.Hour)
	jwtToken, err := h.JwtUtil.CreateAccessToken(user.ID, exp)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    *jwtToken,
		Expires:  exp,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	return nil
}

//

func (h *Handler) UnprocessableResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, _ = w.Write([]byte(err.Error()))
}

func (h *Handler) LogError(m string, e error) error {
	if e != nil {
		h.Logger.Println(m, e.Error())
	} else {
		h.Logger.Println(e)
	}

	return errors.New(m)
}
