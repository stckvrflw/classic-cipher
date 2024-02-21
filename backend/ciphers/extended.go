package ciphers

func ExtendedVigenereEncrypt(s []byte, keyword string) []byte {
	keyword = generateProperKeyword(len(s), keyword)

	var result []byte
	for i, char := range s {
		// Loop over s as a series of chars
		r := (rune(char) + rune(keyword[i])) % 256
		result = append(result, byte(r))
	}

	return result
}

func ExtendedVigenereDecrypt(s []byte, keyword string) []byte {
	keyword = generateProperKeyword(len(s), keyword)

	var result []byte
	for i, char := range s {
		r := (rune(char) - rune(keyword[i]) + 256) % 256
		result = append(result, byte(r))
	}

	return result
}
