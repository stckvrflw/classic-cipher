package ciphers

// var alphabets = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func VigenereEncrypt(s []byte, keyword string) []byte {
	keyword = generateProperKeyword(len(s), keyword)

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

func VigenereDecrypt(s []byte, keyword string) []byte {
	keyword = generateProperKeyword(len(s), keyword)

	var result []byte
	for i, char := range s {
		if isAlpha(rune(char)) {
			r := (rune(char) - rune(keyword[i]) + 26) % 26
			r += 'A'
			result = append(result, byte(r))
		} else {
			result = append(result, byte(char))
		}
	}

	return result
}
