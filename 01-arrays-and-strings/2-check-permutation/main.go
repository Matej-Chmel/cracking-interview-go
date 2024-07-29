package main

import (
	"sort"

	goi "github.com/Matej-Chmel/go-interview"
)

type RuneSlice []rune

func (s RuneSlice) Len() int {
	return len(s)
}

func (s RuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s RuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func sorted(s string) string {
	slice := RuneSlice([]rune(s))
	sort.Sort(slice)
	return string(slice)
}

func bySorting(a, b string) bool {
	a, b = sorted(a), sorted(b)
	return a == b
}

func byCounting(a, b string) bool {
	chars := [128]int{}

	for _, v := range a {
		chars[v]++
	}

	for _, v := range b {
		if chars[v]--; chars[v] < 0 {
			return false
		}
	}

	for _, v := range chars {
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	it := goi.NewInterview2[string, string, bool]()

	it.AddCase("abc", "bca", true)
	it.AddCase("abc AB", "AB abc", true)
	it.AddCase("hello", "llohe", true)
	it.AddCase("hello", "world", false)
	it.AddCase("abc", "bc", false)
	it.AddCase("", " ", false)

	it.AddSolutions(bySorting, byCounting)
	it.Print()
}
