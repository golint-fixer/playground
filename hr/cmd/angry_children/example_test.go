package main

import (
	"fmt"
	"log"
)

func ExampleMinUnfair() {
	// Calculate the minimum unfairness.
	k := 3
	pkts := []int{10, 100, 300, 200, 1000, 20, 30}
	min, err := MinUnfair(pkts, k)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(min)
	// Output: 20
}
