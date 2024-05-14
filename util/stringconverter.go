package util

import (
	"strings"
)

func ExtractSpecificCharacters(str string, spchars []string) []string {
	chars := strings.Split(str, "")
	var modifiedStrings []string
	for _, c := range chars {
		if CompareStrings(c, spchars) {
			modifiedStrings = append(modifiedStrings, c)
		}
	}
	return modifiedStrings
}

func CompareStrings(target string, compareList []string) bool {
	for _, compareItem := range compareList {
		if target == compareItem {
			return true
		}
	}
	return false
}