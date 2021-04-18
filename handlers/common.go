package handlers

import (
	"html/template"
	"net/http"
	"time"

	"../services"
)

const PartialsPath = "./templates/partials/"
const ViewsPath = "./templates/views/"
const LayoutName = "layout.gohtml"

var Cache services.Cache
var SessionDuration time.Duration
var SessionTokenLength int

func render(name string, data interface{}, w http.ResponseWriter){
	templateFile := name + ".gohtml"
	page, err := template.ParseFiles(ViewsPath + templateFile,PartialsPath+LayoutName)
	page = template.Must(page,err)
	page.ExecuteTemplate(w,"layout", data)
}
