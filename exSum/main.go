package main

import "fmt"

func main() {
	var n uint64 = 4
	fmt.Println(expSum(n))
}

func expSum(n uint64) uint64 {
	partitions := make([]uint64, n+1)
	partitions[0] = 1
	var i uint64 = 1
	for ; i <= n; i++ {
		j := i
		for ; j <= n; j++ {
			partitions[j] += partitions[j-i]
		}
	}

	return partitions[n]
}
