package ciphers

func AffineEncrypt(s []byte, A, B int) []byte {
	var result []byte
	for _, char := range s {
		result = append(result, byte((A*int(char-'A')+B)%26+'A'))
	}

	return result
}

func AffineDecrypt(s []byte, A, B int) []byte {
	var result []byte
	A_inverse := modInverse(A, 26)
	for _, char := range s {
		// (((a_inv * ((cipher[i]+'A' - b)) % 26)) + 'A');
		// result = append(result, byte((A_inverse*(int(char+'A')-B)%26)))
		result = append(result, byte(((A_inverse)*(int(char+'A')-B)%26)+'A'))
	}

	return result
}
