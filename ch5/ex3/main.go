package main

import "fmt"

func prefixer(s string) func(string) string {
	return func(p string) string {
		return fmt.Sprintf("%s %s", s, p)
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
