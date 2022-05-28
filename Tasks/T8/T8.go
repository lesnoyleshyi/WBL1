package main

import (
	"fmt"
	"log"
)

func main() {
	var num int64 = 100

	changedNum, ok := changeBit(num, 5, 0)
	if !ok {
		log.Fatal("incorrect usage")
	}
	changedNum2, ok := changeBit(num, 5, 1)
	if !ok {
		log.Fatal("incorrect usage")
	}

	fmt.Println("original number:", num)
	fmt.Println("changed number:", changedNum)
	fmt.Println("backward changed number:", changedNum2)

}

//changeBit sets "i" bit to 0 or 1 (val param).
//It returns 0, false on failure (wrong "i" or "val").
func changeBit(num int64, i int8, val int8) (int64, bool) {
	if i > 64 {
		return 0, false
	}
	if val == 1 {
		return num | (1 << i), true
	} else if val == 0 {
		return num & ^(1 << i), true
	} else {
		return 0, false
	}
}
