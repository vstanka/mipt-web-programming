package main

import (
	"strings"
	"unicode"
)

func RemoveEven(slice []int) (noEvenSlice []int) {
	for _, v := range slice {
		if v & 1 == 1 {
			noEvenSlice = append(noEvenSlice, v)
		}
	}
	return
}

func PowerGenerator(s int) func() int {
	var mantisa, curnum = s, 1
	return func() int {
		curnum *= mantisa
		return curnum
	}
}

func DifferentWordsCount(s string) (count int) {
	runeChecker := func (r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return '#'
	}
	s = strings.Map(runeChecker, s)

	words := strings.Split(s,"#")
	differWords := make(map[string]bool)

	for _, v := range words {
		if v != "" {
			differWords[v] = true
		}
	}
	return len(differWords)
}


