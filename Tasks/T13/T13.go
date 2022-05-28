package main

import "fmt"

func main() {
	n1, n2 := 1, 2
	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}
