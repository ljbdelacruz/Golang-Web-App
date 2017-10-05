package my_strings

func Splitter(item string, exclude string)(string){
	data:="";
	x:=0;
	for x < len(item) {
		//this is how to access each character of a string
		yo := string([]rune(item)[x]);
		if yo != exclude {
			data+=yo;
		}
		x+=1;
	}
	return data;
}

