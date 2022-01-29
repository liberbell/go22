package main

import "fmt"

var env string

func init() {
	fmt.Println("Secret key from package main fetched.")
}

func main() {
	aa
}
