package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println(SqCubRevPrime(100))
	fmt.Println(89*89)
	fmt.Println(89*89*89)
	
}

func SqCubRevPrime(n int) int {
	counter := 0
	res := 89    
	for counter != n {
		if big.NewInt(int64(revNum(res*res))).ProbablyPrime(20) && big.NewInt(int64(revNum(res*res*res))).ProbablyPrime(20) {
			counter++
		}
		res++
	}
	return res - 1
}

func revNum(n int) int {
	nbr := 0
	for n != 0 {
		digit := n % 10
		nbr = nbr*10 + digit
		n /= 10
	}
	return nbr
}
