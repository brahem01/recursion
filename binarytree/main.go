package main

import (
	"fmt"
)

func main() {
	recu("", 0, 0)
}

func recu(s string, sum, sum2 int)  {
	if len(s)== 4 {
		if sum == sum2 {
			fmt.Println(s)
		}
		return
	}
	recu(s+"0", sum, sum2)
	if len(s)< 2 {
		recu(s+"1", sum+1, sum2)
	}else{
		recu(s+"1", sum, sum2+1)
	}
}







