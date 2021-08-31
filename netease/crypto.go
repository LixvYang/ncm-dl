package netease

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"time"
)

func Encrypt(origData []byte) (params,encSecKey string,err error) {
	enc1, err := aesCBCEncrypt(origData, []byte(PresetKey), []byte(IV))
	if err != nil {
		return
	}

	secKey := createSecretKey(16,Base62) 

}

func createSecretKey(size int,charset string) []byte {
	secKey := make([]byte,size)
	n := len(charset)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range secKey {
		secKey[i] = charset[r.Intn(n)]
	}
	return secKey
}

func aesCBCEncrypt(plainText,secKey,iv []byte) (string,error) {
	block,err := aes.NewCipher(secKey)
	if err != nil {
		return "",err
	}

	plainText = pkcs5Padding(plainText,block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block,iv)
	cipherText := make([]byte,len(plainText))
	blockMode.CryptBlocks(cipherText,plainText)

	// implements base64 encoding as specified
	return base64.StdEncoding.EncodeToString(cipherText),nil
}

func pkcs5Padding(src []byte,blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)},padding)
	return append(src,paddingText...)
}