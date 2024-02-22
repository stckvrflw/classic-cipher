package ciphers

import (
	"strings"
	"unicode"
)

func isAlpha(r rune) bool {
	return unicode.IsLetter(r)
}

func generateProperKeyword(targetLen int, keyword string) string {
	properLen, i := 0, 0
	var result strings.Builder
	for properLen < targetLen {
		result.WriteRune(rune(keyword[i]))
		i = (i + 1) % len(keyword)
		properLen++
	}
	return result.String()
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

func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return -1
}
