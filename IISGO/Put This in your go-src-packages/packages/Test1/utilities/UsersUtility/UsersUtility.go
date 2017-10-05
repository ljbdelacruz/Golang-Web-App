package UsersUtility
import (
	"fmt"
	"net/http"
	"encoding/json"
	"packages/my_json"
//	"packages/general"
	//"image/png"
//	"io"
//	"os"
//	"strings"
//	"encoding/base64"
);
//models
type Users struct{
	ID string `json:"ID"`
	FirstName string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName string `json:"LastName"`
	BirthDate string `json:"BirthDate"`
}
//functionality
func ConvertJsonString(item string)(*Users){
	data := &Users{}
	json.Unmarshal([]byte(item), data);
	return data;
}
//route functionalities
func Insert(w http.ResponseWriter, req *http.Request){
	fmt.Println("Insert New User");
	body:=my_json.ContentLengthToJson(req);
	data:=ConvertJsonString(body);
	fmt.Println(data);
}
func UploadImage(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.Body);
	fmt.Println(req.FormFile("file"));
	// body:=my_json.ContentLengthToJson(req);
	// fmt.Println(body);
	// data:=general.ConvertJsonStringMyImage(body);
	// fmt.Println(data);
	// img, _ := os.Create("image.jpg");
	// defer img.Close()
	// dataurl := strings.Replace(data.Source, "data:image/png;base64,", "", 1);
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(dataurl))
    // //defer reader
	// b, _ := io.Copy(img, reader);
    // fmt.Println("File size: ", b);
}
