package main

import (
	"backend/ciphers"
	"fmt"
)

func vigenere_test() {
	str := []byte("ABCD E")
	key := "LALA"
	fmt.Println(string(str))
	str = ciphers.Sanitize(str)
	fmt.Println(string(str))
	encrypt := ciphers.VigenereEncrypt(str, key)
	fmt.Println(string(encrypt))

	decrypt := ciphers.VigenereDecrypt(encrypt, key)
	fmt.Println(string(decrypt))

	// b64Orig := base64.StdEncoding.EncodeToString(str)
	// b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	// b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	// fmt.Println(b64Orig)
	// fmt.Println(b64Enc)
	// fmt.Println(b64Dec)
}

func main() {
	vigenere_test()
}
