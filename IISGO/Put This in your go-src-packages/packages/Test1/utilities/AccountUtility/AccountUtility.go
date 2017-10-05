package AccountUtility
import (
	"fmt"
	"net/http"
	"encoding/json"
	"packages/my_json");
//Models
type Account struct{
	Username string `json:"username"`
	Password string `json:"password"`
}
type JsonResponse struct{
	Success bool `json:"Success"`
	//this is how to make json data vague not specific type so you can put any json data inside it
	//it will be treated as an object instead of specific values
	Data interface{} `json:"Data"`
}
//functionalities
func ConvertJsonString(item string)(*Account){
	data := &Account{}
	json.Unmarshal([]byte(item), data);
	return data;
}
//posting data
//check username and password
func CheckAuthentication(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Check Username and Password");
	body:=my_json.ContentLengthToJson(req);
	data:=ConvertJsonString(body);
	fmt.Println(data);
	acct:=Account{Username:"Naniii"};
	result:=false;
	if(data.Username == "admin" && data.Password == "admin"){
		result=true;
	}
	u := JsonResponse{Success: result, Data: acct};
	fmt.Println(u);
	json.NewEncoder(w).Encode(u);
}

