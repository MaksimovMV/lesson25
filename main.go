package main

import (
	"flag"
	"fmt"
	"strings"
)

func stringSearching(str string, substr string) bool {
	strRune := []rune(str)
	substrRune := []rune(substr)

	for i := 0; i <= len(strRune)-len(substrRune); i++ {
		for j := 0; substrRune[j] == strRune[j+i]; j++ {
			if j == len(substrRune)-1 {
				return true
			}
		}
	}

	return false
}

func main() {
	var str, substr string

	flag.StringVar(&str, "str", "default", "set string")
	flag.StringVar(&substr, "substr", "default", "set substring")

	flag.Parse()

	fmt.Println(stringSearching(str, substr))

	fmt.Println(strings.Contains(str, substr))
}
