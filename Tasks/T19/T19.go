package main

import (
	"fmt"
	"log"
)

func main() {
	var input string

	fmt.Println("Enter string value:")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatalln(err)
	}
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()
}
