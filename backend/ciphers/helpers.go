package ciphers

import "unicode"

func isAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

// strip whitespaces, numbers, and punctuations, and convert to UPPERCASE
func Sanitize(s []byte) []byte {
	var result []byte
	for _, b := range s {
		if isAlpha(rune(b)) {
			if unicode.IsLower(rune(b)) {
				b = byte(unicode.ToUpper(rune(b)))
			}
			result = append(result, b)
		}
	}
	return result
}
