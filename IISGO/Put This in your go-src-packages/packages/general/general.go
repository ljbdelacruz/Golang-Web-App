package general
import (
	"fmt"
	"encoding/json"
//	"packages/my_strings"
)
type MyImage struct{
	Source string `json:"Source"`
}
func ConvertJsonStringMyImage(item string)(*MyImage){
	data := &MyImage{}
	fmt.Println("------");
	fmt.Println(item);
	json.Unmarshal([]byte(item), &data);
	return data;
}