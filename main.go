// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input_str := os.Args[1:]
	count := WordCountept(input_str[0])
	fmt.Println(count)
}

func WordCountept(src string) int {
	srez := strings.Fields(src)
	return len(srez)
}
