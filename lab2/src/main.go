package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var alphabet = []rune{'a', 'b', 'c'}

const maxLen = 20
const numTests = 1000

func makeRandomWord() string {
	var b strings.Builder
	length := rand.Intn(maxLen + 1)
	for range length {
		b.WriteRune(alphabet[rand.Intn(len(alphabet))])
	}
	return b.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	nonDet := buildNFA()
	alt := buildAFA()

	everythingOK := true

	for range numTests {
		w := makeRandomWord()

		r := matchByRegex(w)
		d := DFA(w)
		n := NFA(nonDet, w)
		a := AFA(alt, w)

		if !((r == d) && (r == n) && (r == a)) {
			everythingOK = false
			fmt.Printf("Слово не прошло по всем распознавателям %q:\n", w)
			fmt.Printf("  regex: %v\n", r)
			fmt.Printf("  DFA  : %v\n", d)
			fmt.Printf("  NFA  : %v\n", n)
			fmt.Printf("  AFA  : %v\n", a)
			break
		}
	}

	fmt.Println()
	fmt.Println("Все тесты пройдены", everythingOK)
}
