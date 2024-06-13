package main

import (
	"strings"

	goi "github.com/Matej-Chmel/go-interview"
)

type A interface {
	AddCase([]rune, []rune)
}

func addCase(i A, s string) {
	spaces := strings.Count(s, " ") * 2
	data := s + strings.Repeat(" ", spaces)
	solved := strings.ReplaceAll(s, " ", "%20")
	i.AddCase([]rune(data), []rune(solved))
}

func twoPointers(slice []rune) []rune {
	l := len(slice)

	for l >= 0 {
		if l--; slice[l] != ' ' {
			break
		}
	}

	r := len(slice) - 1

	for l >= 0 {
		if slice[l] == ' ' {
			slice[r-2] = '%'
			slice[r-1] = '2'
			slice[r] = '0'
			r -= 3
		} else {
			slice[r] = slice[l]
			r--
		}

		l--
	}

	return slice
}

func main() {
	iv := goi.NewInterview[[]rune, []rune]()
	i := &iv

	addCase(i, "Mr John Smith")
	addCase(i, "Hello World")
	addCase(i, "... ...")
	addCase(i, "protocol://data   data .com")

	i.AddSolution(twoPointers)
	i.Print()
}
