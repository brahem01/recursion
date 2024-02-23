package main

import "fmt"

func main() {
	strOfLengthK([]string{"a", "b"}, 3, "")
}


func strOfLengthK(slice []string, n int, s string)  {
	if len(s) == n {
		fmt.Println(s)
		return 
	}
	for _, v:= range slice {
		strOfLengthK(slice, n, s+v)
	}
}