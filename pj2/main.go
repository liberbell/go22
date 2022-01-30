package main

import "fmt"

var env string

func init() {
	fmt.Println("Secret key from package main fetched.")
}

func init() {
	fmt.Println("Database initialized and ready.")
}

func init() {
	env = "development"
}

func main() {
	fmt.Println(env + " environment loaded.")
}
