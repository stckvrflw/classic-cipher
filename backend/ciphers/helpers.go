package ciphers

import "unicode"

func isAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

// strip whitespaces, numbers, and punctuations
func Sanitize(s []byte) []byte {
	var result []byte
	for _, b := range s {
		if isAlpha(rune(b)) {
			result = append(result, b)
		}
	}
	return result
}
