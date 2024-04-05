package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AES + CBC(Cipher Block Chaining) + PKCS7Padding
// echo -n 'plainText' | openssl aes-128-cbc -e -K 'hex(key)' -iv 'hex(iv)' -nosalt | hexdump -e '16/1 "%02x"'
func AESCBCEncrypt(plainText, key []byte) (ciphertext []byte, err error) {
	plainText = PKCS7Padding(plainText, aes.BlockSize)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, nil
}

func AESCBCDecrypt(ciphertext, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return PKCS7UnPadding(ciphertext), nil
}

// TODO recommand: AES GCM
