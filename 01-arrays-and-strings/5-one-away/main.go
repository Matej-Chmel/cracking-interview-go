package main

import (
	goi "github.com/Matej-Chmel/go-interview"
)

func shortLong(a, b string) (string, string) {
	if len(a) < len(b) {
		return a, b
	}

	return b, a
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

func twoMethods(a, b string) bool {
	if len(a) == len(b) {
		return editReplace(a, b)
	} else if len(a)-1 == len(b) {
		return editInsert(b, a)
	} else if len(a) == len(b)-1 {
		return editInsert(a, b)
	}

	return false
}

func unifiedMethod(a, b string) bool {
	shorter, longer := shortLong(a, b)
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
	i := goi.NewInterview2[string, string, bool]()

	i.AddCase("pale", "ple", true)
	i.AddCase("pales", "pale", true)
	i.AddCase("pale", "bale", true)
	i.AddCase("pale", "bae", false)

	i.AddSolutions(twoMethods, unifiedMethod)
	i.Print()
}
