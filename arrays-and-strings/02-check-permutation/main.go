package main

import (
	"sort"

	goi "github.com/Matej-Chmel/go-interview"
)

type Pair struct {
	a string
	b string
}

func sorted(s string) string {
	slice := []rune(s)

	sort.Slice(slice, func(i int, j int) bool {
		return slice[i] < slice[j]
	})

	return string(slice)
}

func bySorting(p Pair) bool {
	a := sorted(p.a)
	b := sorted(p.b)
	return a == b
}

func main() {
	it := goi.NewInterview[Pair, bool]()

	it.AddCase(Pair{"abc", "bca"}, true)
	it.AddCase(Pair{"abc AB", "AB abc"}, true)
	it.AddCase(Pair{"hello", "llohe"}, true)
	it.AddCase(Pair{"hello", "world"}, false)
	it.AddCase(Pair{"abc", "bc"}, false)
	it.AddCase(Pair{"", " "}, false)

	it.AddSolution(bySorting)
	it.Print()
}
