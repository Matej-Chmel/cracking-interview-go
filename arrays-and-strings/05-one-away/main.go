package main

import (
	goi "github.com/Matej-Chmel/go-interview"
)

type Pair struct {
	A string
	B string
}

func (p *Pair) shortLong() (string, string) {
	if len(p.A) < len(p.B) {
		return p.A, p.B
	}

	return p.B, p.A
}

func editReplace(shorter string, longer string) bool {
	s, l := 0, 0
	diff := false

	for s < len(shorter) && l < len(longer) {
		if shorter[s] != longer[l] {
			if diff {
				return false
			}

			diff = true
		}

		l++
		s++
	}

	return true
}

func editInsert(shorter string, longer string) bool {
	s, l := 0, 0

	for s < len(shorter) && l < len(longer) {
		if shorter[s] != longer[l] {
			if s != l {
				return false
			}

			l++
		} else {
			l++
			s++
		}
	}

	return true
}

func twoMethods(p Pair) bool {
	a, b := p.A, p.B

	if len(a) == len(b) {
		return editReplace(a, b)
	} else if len(a)-1 == len(b) {
		return editInsert(b, a)
	} else if len(a) == len(b)-1 {
		return editInsert(a, b)
	}

	return false
}

func unifiedMethod(p Pair) bool {
	shorter, longer := p.shortLong()
	s, l := 0, 0
	diff := false

	for s < len(shorter) && l < len(longer) {
		if shorter[s] == longer[l] {
			s++
		} else {
			if diff {
				return false
			}

			diff = true

			if len(shorter) == len(longer) {
				s++
			}
		}

		l++
	}

	return true
}

func main() {
	i := goi.NewInterview[Pair, bool]()

	i.AddCase(Pair{"pale", "ple"}, true)
	i.AddCase(Pair{"pales", "pale"}, true)
	i.AddCase(Pair{"pale", "bale"}, true)
	i.AddCase(Pair{"pale", "bae"}, false)

	i.AddSolutions(twoMethods, unifiedMethod)
	i.Print()
}
