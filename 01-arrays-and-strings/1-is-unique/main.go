package main

import (
	goi "github.com/Matej-Chmel/go-interview"
)

func withArray(s string) bool {
	a := [128]bool{}

	for _, c := range s {
		if a[c] {
			return false
		}

		a[c] = true
	}

	return true
}

func withLoop(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	it := goi.NewInterview[string, bool]()

	it.AddCase("abc", true)
	it.AddCase("abcABC", true)
	it.AddCase("abc def", true)
	it.AddCase("abca", false)
	it.AddCase("AbcAb", false)
	it.AddCase("abc abc A", false)

	it.AddSolutions(withArray, withLoop)
	it.Print()
}
