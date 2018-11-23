package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func main() {
	wc.Test(WordCount)
	WordCount("I am kene")
}

func WordCount(s string) map[string]int  {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
		//if _, ok := m[v]; ok == true {
		//	m[v] += 1
		//} else {
		//	m[v] = 1
		//}
	}

	fmt.Print(strings.Fields(s))
	return m
}