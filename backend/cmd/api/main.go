package main

import (
	"backend/ciphers"
	"encoding/base64"
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

	fmt.Println("\nBASE64")

	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println(b64Orig)
	fmt.Println(b64Enc)
	fmt.Println(b64Dec)
}

func auto_vigenere_test() {
	str := []byte("qwertyUIOPASDFGHJKLZXCVBNM")
	fmt.Println(string(str))
	str = ciphers.Sanitize(str)
	fmt.Println(string(str))

	encrypt := ciphers.AutoVigenereEncrypt(str)
	fmt.Println(string(encrypt))

	decrypt := ciphers.AutoVigenereDecrypt(encrypt)
	fmt.Println(string(decrypt))

	fmt.Println("\nBASE64")

	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println(b64Orig)
	fmt.Println(b64Enc)
	fmt.Println(b64Dec)
}

func extended_vigenere_test() {
	str := []byte("qwertyUIOP ASDFGHJKL ZXCVBNM 12#45^&")
	key := "HELLOHIWHATISUP"
	fmt.Println(string(str))

	encrypt := ciphers.ExtendedVigenereEncrypt(str, key)
	fmt.Println(string(encrypt))

	decrypt := ciphers.ExtendedVigenereDecrypt(encrypt, key)
	fmt.Println(string(decrypt))

	fmt.Println("\nBASE64")

	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println(b64Orig)
	fmt.Println(b64Enc)
	fmt.Println(b64Dec)
}

func playfair_test() {
	str := []byte("instrumentsz")
	key := "MONARCHY"
	fmt.Println(string(str))
	str = ciphers.Sanitize(str)
	fmt.Println(string(str))

	encrypt := ciphers.PlayfairEncrypt(str, key)
	fmt.Println(string(encrypt))

	decrypt := ciphers.PlayfairDecrypt(encrypt, key)
	fmt.Println(string(decrypt))

	// fmt.Println("\nBASE64")

	// b64Orig := base64.StdEncoding.EncodeToString(str)
	// b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	// b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	// fmt.Println(b64Orig)
	// fmt.Println(b64Enc)
	// fmt.Println(b64Dec)
}

func main() {
	// vigenere_test()
	// auto_vigenere_test()
	// extended_vigenere_test()
	playfair_test()
}
