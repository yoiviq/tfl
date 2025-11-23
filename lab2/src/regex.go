package main

import (
	"regexp"
)

func matchByRegex(word string) bool {
	regexPattern := regexp.MustCompile(`^(?:abc|bac|ac|bc)*ba[abc][ab][bc]$`)
	return regexPattern.MatchString(word)
}