package main
import (
	"fmt"
)


func main(){
	fmt.Printf("Testo\n");
	list:=[6]int{1,2,3,4,5,6};
	
	for(i=0;i<cap(list);i++){
		fmt.Printf(list[i]);
	}

}