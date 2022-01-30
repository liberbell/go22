package main

import "fmt"

func main() {
	anon := func(msg string) bool {
		fmt.Println(msg)
		return true
	}
}
