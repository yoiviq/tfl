package main

import (
	"fmt"
)

func countRunes(w string, r rune) int {
	count := 0
	for _, ch := range w {
		if ch == r {
			count++
		}
	}
	return count
}

func pCount(w string) int {
	return countRunes(w, 'p')
}

func bCount(w string) int {
	return countRunes(w, 'b')
}

func modAB(w string) int {
	return (countRunes(w, 'a') + countRunes(w, 'b')) % 3
}

func checkInvariants(chain []string) (ok1, ok2, ok3 bool) {
	if len(chain) == 0 {
		return true, true, true
	}

	ref1 := pCount(chain[0])
	ref2 := bCount(chain[0])
	ref3 := modAB(chain[0])

	ok1, ok2, ok3 = true, true, true

	for _, w := range chain[1:] {
		if pCount(w) > ref1 {
			ok1 = false
		}
		if bCount(w) < ref2 {
			ok2 = false
		}
		if modAB(w) != ref3 {
			ok3 = false
		}
	}
	return
}

func testInvariants(trials int, showFails int) {
	failCount := 0
	for range trials {
		start := randomWord(10, 25)
		chain := reduce(start, newT, 3, 10)
		ok1, ok2, ok3 := checkInvariants(chain)

		if !(ok1 && ok2 && ok3) {
			failCount++
			if failCount <= showFails {
				fmt.Printf("\n[ОШИБКА #%d]\nНачало: %s\nЦепочка: %v\n", failCount, start, chain)
				fmt.Printf("#p? %v\n", ok1)
				fmt.Printf("#b? %v\n", ok2)
				fmt.Printf("(a+b) mod 3? %v\n", ok3)
			}
		}
	}
	fmt.Printf("\nВсего тестов: %d\nНарушения инвариантов: %d\n", trials, failCount)
}
