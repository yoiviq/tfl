package main

import (
	"fmt"
	"math/rand"
	"strings"
)


type Rule struct {
	pattern, replacement string
}

var (
	T = []Rule{
		{"bpb", "abapba"},
		{"p", "aba"},
		{"apa", "bb"},
		{"abba", "baab"},
	}

	newT = []Rule{
		{"bpb", "abapba"},
		{"p", "aba"},
		{"apa", "bb"},
		{"abba", "baab"},
		{"baababaa", "abbbb"},
		{"aababaab", "bbbba"},
		{"aabaa", "bb"},
	}
)

var alph = []rune{'a', 'b', 'p'}

func randomWord(minLen, maxLen int) string {
	n := rand.Intn(maxLen-minLen+1) + minLen
	var sb strings.Builder
	for range n {
		sb.WriteRune(alph[rand.Intn(len(alph))])
	}
	return sb.String()
}


func rewrite(word string, rules []Rule) []string {
	outcomes := make(map[string]struct{})
	for _, rule := range rules {
		searchFrom := 0
		for {
			pos := strings.Index(word[searchFrom:], rule.pattern)
			if pos == -1 {
				break
			}
			pos += searchFrom
			newWord := word[:pos] + rule.replacement + word[pos+len(rule.pattern):]
			outcomes[newWord] = struct{}{}
			searchFrom = pos + 1
		}
	}
	result := make([]string, 0, len(outcomes))
	for w := range outcomes {
		if w != word {
			result = append(result, w)
		}
	}
	return result
}


func reduce(start string, rules []Rule, minSteps, maxSteps int) []string {
	current := start
	trace := []string{current}
	steps := rand.Intn(maxSteps-minSteps+1) + minSteps

	for range steps {
		nextWords := rewrite(current, rules)
		if len(nextWords) == 0 {
			break
		}
		current = nextWords[rand.Intn(len(nextWords))]
		trace = append(trace, current)
	}
	return trace
}

type state struct {
	word  string
	depth int
}

func isReachable(src, dst string, rules []Rule, maxDepth, maxNodes int) (bool, []string) {
	if src == dst {
		return true, []string{src}
	}

	queue := []state{{src, 0}}
	prev := map[string]string{src: ""}
	visited := map[string]bool{src: true}
	nodeCount := 0

	for len(queue) > 0 && nodeCount < maxNodes {
		current := queue[0]
		queue = queue[1:]
		nodeCount++

		if current.depth >= maxDepth {
			continue
		}
		for _, next := range rewrite(current.word, rules) {
			if visited[next] {
				continue
			}
			visited[next] = true
			prev[next] = current.word

			if next == dst {
				path := []string{}
				for w := next; w != ""; w = prev[w] {
					path = append([]string{w}, path...)
				}
				return true, path
			}
			queue = append(queue, state{next, current.depth + 1})
		}
	}
	return false, nil
}

func fuzzTest(trials int, showFailures int) {
	var success, fail, shown int

	for range trials {
		start := randomWord(10, 25)
		seq := reduce(start, T, 1, 8)
		target := seq[len(seq)-1]

		reachable, _ := isReachable(start, target, newT, 35, 1_000_000)
		if reachable {
			success++
		} else {
			fail++
			if shown < showFailures {
				shown++
				fmt.Printf("\nОшибка #%d\n", shown)
				fmt.Printf("Старт: %s\nКонец:   %s\nЦепочка: %v\n", start, target, seq)
			}
		}
	}
	fmt.Printf("Кол-во тестов: %d\nУспех: %d\nОшибки: %d\n", success+fail, success, fail)
}

func main() {
	fuzzTest(1000, 5)
}
