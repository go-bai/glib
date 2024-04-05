## example

```golang
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/go-bai/glib/crypto"
)

func main() {
	plaintext := []byte("hello world")
	key := []byte("1234567812345678")
	ciphertext, err := crypto.AESCBCEncrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	ciphertextHex := hex.EncodeToString(ciphertext)
	fmt.Printf("ciphertextHex: %s\n", ciphertextHex)

	p, err := crypto.AESCBCDecrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plaintext: \"%s\"\n", string(p))
}
```

## reference

- https://go.dev/src/crypto/cipher/example_test.go
- [AES-CBC密文填充攻击—深入理解和编程实现](https://www.packetmania.net/2020/12/01/AES-CBC-PaddingOracleAttack/)
- [【密码学】为什么不推荐在对称加密中使用CBC工作模式](https://cloud.tencent.com/developer/article/2290733)