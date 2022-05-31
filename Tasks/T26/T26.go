package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	set := make(map[rune]struct{})

	rns := []rune(strings.ToLower(s))
	for _, rn := range rns {
		set[rn] = struct{}{}
	}

	if len(set) == len(rns) {
		return true
	}

	return false
}

func main() {
	s1 := "abcd"
	s2 := "abCdefA"
	s3 := "aabcd"

	fmt.Printf("'%s' - %t\n", s1, isUnique(s1))
	fmt.Printf("'%s' - %t\n", s2, isUnique(s2))
	fmt.Printf("'%s' - %t\n", s3, isUnique(s3))
}
