package main

import "fmt"

func main() {
	var s string = "string"
	var n int = 10
	var t bool = true
	ch := make(chan struct{})
	var n64 int64 = 10

	fmt.Println(whatIsIt(s))
	fmt.Println(whatIsIt(n))
	fmt.Println(whatIsIt(t))
	fmt.Println(whatIsIt(ch))
	fmt.Println(whatIsIt(n64))
}

func whatIsIt(smth interface{}) string {
	switch smth.(type) {
	case string:
		return fmt.Sprintf("It's string")
	case int:
		return fmt.Sprintf("It's int")
	case bool:
		return fmt.Sprintf("It's bool")
	case chan struct{}:
		return fmt.Sprintf("It's chan")
	default:
		return fmt.Sprintf("It's smth unknown")
	}
}
