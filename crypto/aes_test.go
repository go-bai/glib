package crypto

import (
	"encoding/hex"
	"testing"
)

func TestAESCBCEncrypt(t *testing.T) {
	plaintext := []byte("hello world")
	key := []byte("1234567812345678")
	ciphertext, err := AESCBCEncrypt(plaintext, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hex.EncodeToString(ciphertext))
}

func TestAESCBCDecrypt(t *testing.T) {

	ciphertextHex := "d4149dcbdd313d0a9e5505876779a18c04e1c775f3a292aab18f848547fdedcd"
	ciphertext, _ := hex.DecodeString(ciphertextHex)
	key := []byte("1234567812345678")
	plaintext, err := AESCBCDecrypt(ciphertext, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\"%s\"", string(plaintext))
}
