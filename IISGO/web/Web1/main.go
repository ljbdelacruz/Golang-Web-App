package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("layouts/*"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	http.ListenAndServe(":8080", mux)
}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}
