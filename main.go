package main
import (
	"log"
	"os"
	"text/template")


var tpl *template.Template


func init(){
	tpl=template.Must(template.ParseFiles("index.html"))
}

type person struct {
	Name string
	Age string
}

func main(){
	models:=[]person{{Name:"Buddha", Age:"13"}, {Name:"John", Age:"90"}}

	err:=tpl.ExecuteTemplate(os.Stdout, "index.html", models)
	if err != nil {
		log.Fatalln(err);
	}
}




