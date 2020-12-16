package libraries

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
)

func AesCFBEncrypt(data []byte, key []byte, iv []byte) []byte {
	b, _ := aes.NewCipher(key)
	if len(iv) < 16 {
		_m := md5.Sum(iv)
		iv = make([]byte, 16)
		for k, v := range _m[:8] {
			iv[k*2] = hextable[v>>4]
			iv[k*2+1] = hextable[v&0x0f]
		}
	}
	s := cipher.NewCFBEncrypter(b, iv)
	dst := make([]byte, len(data))
	s.XORKeyStream(dst, data)
	return dst
}

func AesCFBDecrypt(data []byte, key []byte, iv []byte) []byte {
	a, _ := aes.NewCipher(key)
	if len(iv) < 16 {
		_m := md5.Sum(iv)
		iv = make([]byte, 16)
		for k, v := range _m[:8] {
			iv[k*2] = hextable[v>>4]
			iv[k*2+1] = hextable[v&0x0f]
		}
	}
	s := cipher.NewCFBDecrypter(a, iv)
	dst := make([]byte, len(data))
	s.XORKeyStream(dst, data)
	return dst
}
