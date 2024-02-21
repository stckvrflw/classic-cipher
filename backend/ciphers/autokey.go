package ciphers

import (
	"fmt"
	"strings"
)

var salt = "MYNAMEISPETRUS"

func generateProperKeywordAuto(s []byte) string {
	if len(s) < len(salt) {
		return salt
	}

	var result strings.Builder
	result.WriteString(salt)
	i := 0
	for i < len(s)-len(salt) {
		result.WriteByte(s[i])
		i++
	}
	fmt.Println(result.String())
	return result.String()
}

func AutoVigenereEncrypt(s []byte) []byte {
	keyword := generateProperKeywordAuto(s)

	var result []byte
	for i, char := range s {
		// Loop over s as a series of chars
		if isAlpha(rune(char)) {
			r := (rune(char) + rune(keyword[i])) % 26
			r += 'A'
			result = append(result, byte(r))
		} else {
			result = append(result, byte(char))
		}
	}

	return result
}

func AutoVigenereDecrypt(s []byte) []byte {
	var keyword strings.Builder
	keyword.WriteString(salt)

	var result []byte
	for i, char := range s {
		if isAlpha(rune(char)) {
			r := (rune(char) - rune(keyword.String()[i]) + 26) % 26
			r += 'A'
			result = append(result, byte(r))
			keyword.WriteRune(r)
		} else {
			result = append(result, byte(char))
		}
	}
	return result
}
