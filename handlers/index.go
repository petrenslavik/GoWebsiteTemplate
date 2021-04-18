package handlers

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}
	render("index", nil, w)
}
