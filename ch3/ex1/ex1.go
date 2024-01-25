package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstSlice := greetings[:2]
	secondSlice := greetings[1:4]
	thirdSlice := greetings[3:]
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)
	fmt.Println(thirdSlice)
}
