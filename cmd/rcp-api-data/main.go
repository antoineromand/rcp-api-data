package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}