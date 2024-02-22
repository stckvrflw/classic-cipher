package main

import (
	"backend/ciphers"
	"encoding/base64"
	"fmt"
)

func vigenere_test() {
	fmt.Println("\n-----VIGENERE CIPHER TEST-----")
	str := []byte("ABCD E")
	key := "LALA"
	fmt.Println("ORIGINAL STRING: ", string(str))
	fmt.Println("KEY: ", key)

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.VigenereEncrypt(str, key)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	decrypt := ciphers.VigenereDecrypt(encrypt, key)
	fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	fmt.Println("DECIPHERED base64: ", b64Dec)
}

func auto_vigenere_test() {
	fmt.Println("\n-----AUTOKEY VIGENERE CIPHER TEST-----")
	str := []byte("qwertyUIOPASDFGHJKLZXCVBNM")
	fmt.Println("ORIGINAL STRING: ", string(str))

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.AutoVigenereEncrypt(str)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	decrypt := ciphers.AutoVigenereDecrypt(encrypt)
	fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	fmt.Println("DECIPHERED base64: ", b64Dec)
}

func extended_vigenere_test() {
	fmt.Println("\n-----EXTENDED VIGENERE CIPHER TEST-----")
	str := []byte("qwertyUIOP ASDFGHJKL ZXCVBNM 12#45^&")
	key := "HELLOHIWHATISUP"

	fmt.Println("ORIGINAL STRING: ", string(str))
	fmt.Println("KEY: ", key)

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.ExtendedVigenereEncrypt(str, key)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	decrypt := ciphers.ExtendedVigenereDecrypt(encrypt, key)
	fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	fmt.Println("DECIPHERED base64: ", b64Dec)
}

func playfair_test() {
	fmt.Println("\n-----PLAYFAIR CIPHER TEST-----")
	str := []byte("instrumentsz")
	key := "MONARCHY"

	fmt.Println("ORIGINAL STRING: ", string(str))
	fmt.Println("KEY: ", key)

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.PlayfairEncrypt(str, key)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	decrypt := ciphers.PlayfairDecrypt(encrypt, key)
	fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	fmt.Println("DECIPHERED base64: ", b64Dec)
}

func affine_test() {
	fmt.Println("\n-----AFFINE CIPHER TEST-----")
	str := []byte("HELLO")
	A := 5
	B := 8

	fmt.Println("ORIGINAL STRING: ", string(str))
	fmt.Printf("A: %d; B: %d\n", A, B)

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.AffineEncrypt(str, A, B)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	decrypt := ciphers.AffineDecrypt(encrypt, A, B)
	fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	fmt.Println("DECIPHERED base64: ", b64Dec)
}

func hill_test() {
	fmt.Println("\n-----HILL CIPHER TEST-----")
	str := []byte("ACT")
	key := "GYBNQKURP"
	partition := 3

	fmt.Println("ORIGINAL STRING: ", string(str))
	fmt.Println("KEY: ", key)
	fmt.Println("Partition: ", partition)

	str = ciphers.Sanitize(str)
	fmt.Println("SANITIZED STRING: ", string(str))

	encrypt := ciphers.HillEncrypt(str, key, partition)
	fmt.Println("CIPHERTEXT: ", string(encrypt))

	// decrypt := ciphers.VigenereDecrypt(encrypt, key)
	// fmt.Println("DECIPHERED: ", string(decrypt))

	fmt.Println("\nBASE64")
	b64Orig := base64.StdEncoding.EncodeToString(str)
	b64Enc := base64.StdEncoding.EncodeToString(encrypt)
	// b64Dec := base64.StdEncoding.EncodeToString(decrypt)

	fmt.Println("ORIGINAL base64: ", b64Orig)
	fmt.Println("CIPHERTEXT base64: ", b64Enc)
	// fmt.Println("DECIPHERED base64: ", b64Dec)
}

func main() {
	vigenere_test()
	auto_vigenere_test()
	extended_vigenere_test()
	playfair_test()
	affine_test()
	hill_test()
}
