package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter string value:")
	if !s.Scan() {
		log.Fatal("")
	}

	words := strings.Split(s.Text(), " ")
	fmt.Println(strings.Join(words, " "))
	for i,j := 0, len(words)-1; i < j; i,j = i+1,j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(strings.Join(words, " "))
}
