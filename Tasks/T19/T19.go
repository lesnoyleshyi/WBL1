package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter string value:")
	if !s.Scan() {
		log.Fatal("")
	}
	str := []rune(s.Text())

	for i,j := 0, len(str)-1; i < j; i,j = i+1,j-1 {
		str[i], str[j] = str[j], str[i]
	}

	fmt.Println(string(str))
}
