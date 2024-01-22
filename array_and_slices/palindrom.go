package arrayandslices

import "strings"

func Palindrom(word string) bool {
	
	var newWord string
	for i := len(word) - 1; i >= 0; i-- {
		newWord += strings.ToLower(string(word[i]))
	}
	if strings.ToLower(word) == newWord {
		return true
	}
	return false
}