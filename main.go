// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	vhodnaya_stroka := os.Args[1:]
	count := WordCountept(vhodnaya_stroka[0])
	fmt.Println(count)
}

func WordCountept(src string) int {
	srez := strings.Fields(src)
	var flag int
	for _, a := range srez {
		a = a
		flag++
	}
	return flag
}
