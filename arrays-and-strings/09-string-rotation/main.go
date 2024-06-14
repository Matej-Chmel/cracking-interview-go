package main

import (
	"strings"

	goi "github.com/Matej-Chmel/go-interview"
)

func withLibrary(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return strings.Contains(a+a, b)
}

func withLoop(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	first, i, j := true, 0, 0

	for {
		if a[i] == b[j] {
			j++
		} else {
			j = 0
		}

		i++

		if j == len(b) {
			return true
		}

		if i == len(a) {
			if first {
				first = false
				i = 0
			} else {
				break
			}
		}
	}

	return false
}

func main() {
	i := goi.NewInterview2[string, string, bool]()

	i.AddCase("waterbottle", "erbottlewat", true)
	i.AddCase("hello", "llohe", true)
	i.AddCase("rotation", "tationro", true)
	i.AddCase("abcdef", "defabc", true)
	i.AddCase("string", "trinsg", false)
	i.AddCase("openai", "niaepo", false)
	i.AddCase("abcabc", "bcabca", true)
	i.AddCase("racecar", "carrace", true)

	i.AddSolutions(withLibrary, withLoop)
	i.Print()
}
