package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := []int{}

	for i := 0; i < 100; i++ {
		random = append(random, rand.Intn(101))
	}
	fmt.Println(random, len(random))

	for _, v := range random {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Nevermind")
		}
	}

}
