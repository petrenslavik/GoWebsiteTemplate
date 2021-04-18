package middlewares

import (
	"../services"
	"net/http"
)

var Cache services.Cache

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		username := Cache.Get(sessionToken)
		if username == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
