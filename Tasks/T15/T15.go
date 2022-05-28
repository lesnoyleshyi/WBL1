package main

//Если цель функции - вернуть первые 100 символов, то в случае строки,
//содержащей символы, занимающие больше одного байта (юникод),
//последний символ будет не тем, что ожидается.
//Также длина итоговой строки будет не той, что ожидается.
//Причина в том, что слайс берётся по байтам, а не по рунам.
//
//Присваивание значения глобальной переменной может иметь смысл в некоторых случаях,
//но, как правило, это делает код более сложным и приводит к неявному поведению,
//особенно в конкурентных программах.
//Так как слайс - ссылочный тип, мы можем изменить исходную строку в другой функции.

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	//v := createHugeString(1 << 10)
	v := strings.Repeat("мда_", 1<<9)
	justString = v[:101]
	fmt.Println(justString)
}

func fixedSomeFunc() string {
	v := strings.Repeat("мда_", 1<<9)
	internalString := make([]rune, 100)
	copy(internalString, []rune(v)[:100])
	return string(internalString)
}

func main() {
	someFunc()
	fmt.Println(fixedSomeFunc())
}
