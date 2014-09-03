package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	a := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		a[v]++
	}

	return a
}

func main() {
	wc.Test(WordCount)
}
