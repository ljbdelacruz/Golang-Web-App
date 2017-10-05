package my_json;
import (
	"net/http"
)

func ContentLengthToJson(req *http.Request)(string){
	req.ParseForm();
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs);
	body := string(bs)
	return body;
}

