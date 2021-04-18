package handlers

import (
	"math/rand"
	"net/http"
	"time"
)

var users = map[string]string{
	"TestLogin": "TestPassword",
	"Admin":     "Admin",
}

type Credentials struct {
	Password string
	Username string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var creds Credentials
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		creds.Username = r.Form.Get("username")
		creds.Password = r.Form.Get("password")

		if creds.Username == "" || creds.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expectedPassword, ok := users[creds.Username]

		if !ok || expectedPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionToken := randSeq(SessionTokenLength)
		Cache.Set(sessionToken, creds.Username)

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
			Expires:  time.Now().Add(SessionDuration),
		})

		http.Redirect(w, r, "/home", http.StatusFound)
	} else{
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
