package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password string = "120398"

	fmt.Println("Authorization: Basic", username+":"+password)
}
