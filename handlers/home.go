package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sessionToken, _ := r.Cookie("session_token")
	login := fmt.Sprintf("%v",Cache.Get(sessionToken.Value))
	data := struct{
		Username string
	}{
		login,
	}
	render("home",data,w)
}
