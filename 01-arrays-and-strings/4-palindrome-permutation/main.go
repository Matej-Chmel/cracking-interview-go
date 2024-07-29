package main

import (
	"math/rand"
	"time"

	goi "github.com/Matej-Chmel/go-interview"
)

func randomPermutation(gen *rand.Rand, s string) string {
	runes := []rune(s)

	for i := len(runes) - 1; i > 0; i-- {
		j := gen.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func getValue(r rune) (res int) {
	res = int(r - 'A')

	if res > 25 {
		res = res - 32
	}

	return
}

func withArray(s string) bool {
	a := [26]bool{}

	for _, v := range s {
		idx := getValue(v)

		if idx >= 0 {
			a[idx] = !a[idx]
		}
	}

	odd := false

	for _, v := range a {
		if v && odd {
			return false
		} else if v && !odd {
			odd = true
		}
	}

	return true
}

func toggle(bits *int, i int) {
	mask := 1 << i

	if (*bits & mask) == 0 {
		*bits |= mask
	} else {
		*bits &= ^mask
	}
}

func withBits(s string) bool {
	bits := new(int)

	for _, v := range s {
		val := getValue(v)

		if val >= 0 {
			toggle(bits, val)
		}
	}

	return (*bits & (*bits - 1)) == 0
}

func main() {
	palindromes := [...]string{
		"A man a plan a canal Panama",
		"Madam in Eden Im Adam",
		"Able was I ere I saw Elba",
		"Never odd or even",
		"Was it a car or a cat I saw",
		"Eva can I see bees in a cave",
		"No lemon no melon",
		"Do geese see God",
		"A Santa lived as a devil at NASA",
		"Murder for a jar of red rum",
	}

	notPalindromes := [...]string{
		"The quick brown fox jumps over the lazy dog",
		"Hello world this is a test",
		"Go is a statically typed language",
		"I enjoy solving programming challenges",
		"Artificial intelligence is fascinating",
		"Cloud computing is the future",
		"This sentence is not a palindrome",
		"OpenAI creates advanced AI models",
		"Learning new languages is fun",
		"The weather today is quite pleasant",
	}

	i := goi.NewInterview[string, bool]()
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, v := range palindromes {
		i.AddCase(randomPermutation(gen, v), true)
	}

	for _, v := range notPalindromes {
		i.AddCase(randomPermutation(gen, v), false)
	}

	i.AddSolutions(withArray, withBits)
	i.Print()
}
