package main
import (
	"fmt"
	"html/template"
	"net/http"
)
type hotdog int
func ManageRouteGet(url string, w http.ResponseWriter){
	switch(url){
		case "/":
			tpl.ExecuteTemplate(w, "index.html", "JOHN");
			break;
		case "/ContactUs":
			tpl2.ExecuteTemplate(w, "contactUs.html", "JOHN");
			break;
		default:
			errorTp.ExecuteTemplate(w, "error.html", "JOHN")
			break;
	}
}
func ManageRoutePost(url string, w http.ResponseWriter){
	fmt.Println(url);
	switch(url){
	case "//GetAll":
		fmt.Println("Get All");
		tpl.ExecuteTemplate(w, "index.html", "JOHN");
		break;
	case "/ContactUs":
		tpl2.ExecuteTemplate(w, "contactUs.html", "JOHN");
		break;
	default:
		errorTp.ExecuteTemplate(w, "error.html", "JOHN")
		break;
}

}

//templates
var tpl *template.Template
var tpl2 *template.Template
//template for Error
var errorTp *template.Template

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	if(r.Method=="GET"){
		ManageRouteGet(r.URL.String(), w);
	}else if(r.Method == "POST"){
		if(len(r.Form) > 0){
			fmt.Println(r.Form);
		}
		ManageRoutePost(r.URL.String(), w);
	}
}
func init() {
	//instantiate templates
	tpl = template.Must(template.ParseFiles("./pages/index.html"))
	tpl2 = template.Must(template.ParseFiles("./pages/contactUs.html"))
	errorTp = template.Must(template.ParseFiles("./pages/error.html"))
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

