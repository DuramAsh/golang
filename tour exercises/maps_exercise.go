package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	ss := strings.Fields(s)
	for _, e := range ss {
		m[e]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
