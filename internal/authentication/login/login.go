package login

import (
	"errors"
	"github.com/hugobally/mimiko_api/internal/db/models"
	"net/http"
	"time"
)

// TODO Abuse prevention mecanism
// TODO Does not check cookies -> Will create and log a new account on each endpoint hit, even if the user is logged in

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if h.Config.Env == "DEV" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	}

	if r.Method != http.MethodPost {
		h.InvalidMethod(w)
		return
	}

	user, err := h.InsertUser()
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

func (h *Handler) InsertUser() (*models.User, error) {
	name := "GuestUser"
	u := models.User{Username: &name}
	err := h.Database.Create(&u).Error

	if err != nil {
		return nil, h.LogError("internal server error", err)
	}

	return &u, nil
}

func (h *Handler) SetLoginCookie(w http.ResponseWriter, user *models.User) error {
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

func (h *Handler) InvalidMethod(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	h.LogError("invalid method", nil)
}

func (h *Handler) LogError(m string, e error) error {
	if e != nil {
		h.Logger.Println(m, e.Error())
	} else {
		h.Logger.Println(e)
	}

	return errors.New(m)
}
