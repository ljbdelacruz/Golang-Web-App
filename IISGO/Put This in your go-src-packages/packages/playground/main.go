package main









//sample of converting string json to object
// import (
//     "fmt"
//     "encoding/json"
// )

// type Data struct {
// 	Votes *Votes `json:"votes"`
// 	Username string `json:"username, omitempty"`
// 	Count string `json:"count,omitempty"`
// }
// type Votes struct {
// 	OptionA string `json:"option_A"`
// 	Description string `json:"Description"`
// }

// func main() {
//     s := `{ "votes": { "option_A": "3", "Description":"Somelape" }, "username":"ljbdelacruz" }`
//     data := &Data{
// 		Votes: &Votes{},
//     }
//     err := json.Unmarshal([]byte(s), data)
// 	fmt.Println(err)
// 	fmt.Println(data.Votes.OptionA);
// 	fmt.Println(data.Username);
// 	fmt.Println("WAHT");
//     s2, _ := json.Marshal(data)
//     fmt.Println(string(s2))
//     data.Count = "2"
//     s3, _ := json.Marshal(data)
//     fmt.Println(string(s3))
// }