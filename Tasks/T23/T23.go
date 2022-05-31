package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Remove(elems []string, idx int) []string {
	return append(elems[:idx], elems[idx+1:]...)
	//if order is not critical:
	//elems[idx] = elems[len(elems)-1]
	//return elems[:len(elems)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter some values, separated by space:")
	s, _ := r.ReadString('\n')
	s = strings.Trim(s, "\n")

	fmt.Println("Enter index of elem to delete:")
	var toDel int
	fmt.Scan(&toDel)

	elems := strings.Split(s, " ")
	elems = Remove(elems, toDel)

	fmt.Println("Slice after removal: ", elems)
}
