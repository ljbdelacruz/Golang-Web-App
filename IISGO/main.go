package main
import (
	"packages/Test1/utilities/UsersUtility"
	"packages/Test1/utilities/AccountUtility"
//	"packages/general"
	"html/template"
	//"io"
//	"fmt"
//	 "github.com/julienschmidt/httprouter"
	"net/http"
	// "log"
//	"encoding/json"
	//"strings"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("Layouts/*"))
}
func main() {
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("./Assets/"))));
	//get handler
	//home routes
	http.HandleFunc("/Home/Insert", UsersUtility.Insert);
	http.HandleFunc("/Home/UploadImage", UsersUtility.UploadImage);
	http.HandleFunc("/Home", Home);
	//login routes
	http.HandleFunc("/CheckAuthentication", AccountUtility.CheckAuthentication);
	http.HandleFunc("/", Login);
	http.ListenAndServe(":8080", nil);
}
//templates
func Login(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "Login.html", nil)
}
func Home(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "Home.html", nil)
}
