package crypto

import (
	"bytes"
)

// https://datatracker.ietf.org/doc/html/rfc2315#section-10.3
// pad the data at the trailing end with blockSize - (len(data) mod blockSize) octets
// all having value blockSize - (len(data) mod blockSize).
// 1 < blockSize < 256
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
