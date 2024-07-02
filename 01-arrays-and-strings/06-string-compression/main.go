package main

import (
	"strconv"
	"strings"

	goi "github.com/Matej-Chmel/go-interview"
)

func write(b *strings.Builder, prev byte, count int64) {
	b.WriteByte(prev)
	b.WriteString(strconv.FormatInt(count, 10))
}

func oneLoop(s string) string {
	if len(s) == 0 {
		return s
	}

	var builder strings.Builder
	prev, count := s[0], int64(1)

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			count++
		} else {
			write(&builder, prev, count)
			prev, count = s[i], 1
		}
	}

	write(&builder, prev, count)

	if res := builder.String(); len(res) > len(s) {
		return s
	} else {
		return res
	}
}

func finalLen(s string) (f int) {
	if len(s) == 0 {
		return 0
	}

	prev, count, digits, next10 := s[0], 1, 1, 10

	for i := 1; i < len(s); i++ {
		if prev == s[i] {
			count++

			if count == next10 {
				next10 *= 10
				digits++
			}
		} else {
			f += 1 + digits
			prev, count, digits, next10 = s[i], 1, 1, 10
		}
	}

	f += 1 + digits
	return
}

func countAhead(s string) string {
	f := finalLen(s)

	if len(s) < f || f == 0 {
		return s
	}

	var builder strings.Builder
	builder.Grow(f)
	prev, count := s[0], int64(1)

	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			count++
		} else {
			write(&builder, prev, count)
			prev, count = s[i], 1
		}
	}

	write(&builder, prev, count)
	return builder.String()
}

func getData(s string) (f int, res []int64) {
	if len(s) == 0 {
		return
	}

	prev, digits := s[0], 1
	var count, next10 int64 = 1, 10

	for i := 1; i < len(s); i++ {
		if prev == s[i] {
			count++

			if count == next10 {
				next10 *= 10
				digits++
			}
		} else {
			f += 1 + digits
			res = append(res, int64(prev), count)
			prev, count, digits, next10 = s[i], 1, 1, 10
		}
	}

	f += 1 + digits
	res = append(res, int64(prev), count)
	return
}

func processAhead(s string) string {
	f, data := getData(s)

	if len(s) < f || f == 0 {
		return s
	}

	var builder strings.Builder
	builder.Grow(f)

	for i := 0; i < len(data); i += 2 {
		b, count := byte(data[i]), data[i+1]
		write(&builder, b, count)
	}

	return builder.String()
}

func main() {
	i := goi.NewInterview[string, string]()

	i.AddCase("aabcccccaaa", "a2b1c5a3")
	i.AddCase("AaaAAAaaA", "AaaAAAaaA")
	i.AddCase("AaaAAAaa", "A1a2A3a2")
	i.AddCase("aacczzbbbb", "a2c2z2b4")
	i.AddCase("", "")
	i.AddCase("dddddd", "d6")

	i.AddSolutions(oneLoop, countAhead, processAhead)
	i.Print()
}
