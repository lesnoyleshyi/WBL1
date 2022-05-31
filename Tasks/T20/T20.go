package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	var newStr string

	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		newStr += string(runes[i])
	}
	return newStr
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter string value:")
	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")

	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverse(words[i])
	}

	s = strings.Join(words, " ")
	fmt.Println(s)
}
