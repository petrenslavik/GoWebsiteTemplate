package handlers

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	sessionToken, _ := r.Cookie("session_token")
	Cache.Delete(sessionToken.Value)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Expires:  time.Unix(0,0),
	})

	http.Redirect(w, r, "/", http.StatusFound)
}
