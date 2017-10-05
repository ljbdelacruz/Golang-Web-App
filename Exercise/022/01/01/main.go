package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"log"
)
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("template/*"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/", index);
	mux.GET("/:name", user);
	http.ListenAndServe(":8080", mux);
}
func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"));
}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err:=tpl.ExecuteTemplate(w, "index.html", nil);
	HandleError(w, err);
}
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
