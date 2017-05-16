/**
Hentet fra: https://www.socketloop.com/tutorials/golang-how-to-encrypt-with-aes-crypto
 */
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

//noinspection GoUnusedExportedFunction
func EncryptMessage(message string)[]byte{
	//Her er nøkkelen som begge parter må ha
	block,err := aes.NewCipher([]byte("opensesame123456"))

	if err != nil {
		panic(err)
	}

	str := []byte(message)

	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	//Bruk difi harman her
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)

	return encrypted
}

//noinspection GoUnusedExportedFunction
func DecryptMessage(encrypted []byte) []byte{
	block,err := aes.NewCipher([]byte("opensesame123456"))

	if err != nil {
		panic(err)
	}

	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	decrypter := cipher.NewCFBDecrypter(block, iv) // simple!

	decrypted := make([]byte, len(encrypted))
	decrypter.XORKeyStream(decrypted, encrypted)

	//fmt.Printf("RECIVED: %s\n", decrypted)
	return decrypted
}