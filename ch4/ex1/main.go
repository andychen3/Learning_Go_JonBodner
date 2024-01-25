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

}
