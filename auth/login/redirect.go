package login

import (
	"net/http"
	"net/url"
	"strings"
)

// TODO Build once when starting the app
func (h *Handler) SpotifyRedirect(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("https://accounts.spotify.com")
	if err != nil {
		h.Logger.Println("bad auth url : ", err.Error())
		return
	}
	u.Path += "authorize"

	p := url.Values{}
	p.Add("response_type", "code")
	p.Add("client_id", h.Config.Spotify.ClientId)
	p.Add("redirect_uri", h.Config.Spotify.RedirectUri)
	scopes := []string{
		"streaming",                  // SDK playback

		"user-modify-playback-state", // API playback
		"user-read-playback-state",

		"user-read-email",   // Read Email Address
		"user-read-private", // Subscription type (Free/Premium)

		"playlist-read-private", // Playlist
		"playlist-modify-public",
		"playlist-modify-private",
	}
	p.Add("scope", strings.Join(scopes, " "))

	u.RawQuery = p.Encode()

	http.Redirect(w, r, u.String(), http.StatusFound)
}
