package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

//func main() {
	//Forutsetter at begge parter har samme nøkkel og cipthertest


	//Krypterer input og returnerer den krypterte meldingen
	//message := encryptMessage("Hei, er denne linjen sikker?")

	//Tar den krypterte meldingen som input og skriver ut meldingen som blir sent.
	//decryptMessage(message)
//}

func EncryptMessage(message string)[]byte{
	//Her er nøkkelen som begge parter må ha
	block,err := aes.NewCipher([]byte("opensesame123456"))

	if err != nil {
		panic(err)
	}

	str := []byte(message)

	// 16 bytes for AES-128, 24 bytes for AES-192, 32 bytes for AES-256
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16

	encrypter := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(str))
	encrypter.XORKeyStream(encrypted, str)

	return encrypted
}


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