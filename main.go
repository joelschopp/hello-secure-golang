package main

import (
	"fmt"
	"log"
)

func main() {
	err := SeccompInit()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("hello world")
}
