package login

//
//uuid, err := h.ParseRequest(w, r)
//if err != nil {
//	_, _ = w.Write([]byte(err.Error()))
//	return
//}
//func (h *Handler) ParseRequest(w http.ResponseWriter, r *http.Request) (*string, error) {
//	if r.Method != http.MethodPost {
//		w.WriteHeader(http.StatusMethodNotAllowed)
//		return nil, h.LogError("invalid method", nil)
//	}
//
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		w.WriteHeader(http.StatusUnauthorized)
//		return nil, h.LogError("error while reading body", err)
//	}
//	bodyStr := string(body)
//	if bodyStr == "" {
//		w.WriteHeader(http.StatusUnauthorized)
//		return nil, h.LogError("authentication code not provided", nil)
//	}
//
//	return &bodyStr, nil
//}

//	func (h *Handler) LinkSpotifyAccount(w http.ResponseWriter, r *http.Request) {
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

//func (h *Handler) InsertLinkedAccount(u *spotify.UserResponse, d *spotify.TokenResponse, ctx context.Context) (*models.User, error) {
//
//exp := time.Now().Add(time.Duration(d.ExpiresIn) * time.Second)
//expStr := exp.Format(time.RFC3339)
//
//newUser, err = h.Prisma.CreateLinkedApp(prisma.LinkedAppCreateInput{
//UserId:   u.Id,
//Username: &u.DisplayName,
//Type:     prisma.AppTypeSpotify,
//User: prisma.UserCreateOneWithoutLinkedAppsInput{
//Create: &prisma.UserCreateWithoutLinkedAppsInput{
//Username: &u.DisplayName,
//},
//},
//AccessToken:  &d.AccessToken,
//TokenExpiry:  &expStr,
//RefreshToken: &d.RefreshToken,
//}).User().Exec(ctx)
//}
