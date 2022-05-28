package main

import "fmt"

var strings []string = []string{"cat", "cat", "dog", "cat", "tree", "btree", "man", "dog"}

func main() {
	set := make(map[string]struct{})

	for _, str := range strings {
		set[str] = struct{}{}
	}
	fmt.Println(set)
}
