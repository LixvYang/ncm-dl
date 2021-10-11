package netease

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"ncm-dl/utils"
	"time"
)

const (
	Base62                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PresetKey                   = "0CoJUm6Qyw8W8jud"
	IV                          = "0102030405060708"
	DefaultRSAPublicKeyModulus  = "e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	DefaultRSAPublicKeyExponent = 0x10001
)

func Encrypt(origData []byte) (params,encSecKey string,err error) {
	enc1, err := aesCBCEncrypt(origData, []byte(PresetKey), []byte(IV))
	if err != nil {
		return
	}

	secKey := createSecretKey(16,Base62)

	enc2,err := aesCBCEncrypt([]byte(enc1),secKey,[]byte(IV))
	if err != nil {
		return
	}

	params, encSecKey = enc2, rsaEncrypt(utils.BytesReverse(secKey),DefaultRSAPublicKeyModulus, DefaultRSAPublicKeyExponent)
	return
}

func rsaEncrypt(origData []byte,modulus string,exponent int64) string {
	bigOrigData,bigModulus := new(big.Int),new(big.Int)

	bigOrigData.SetBytes(origData)
	bigModulus.SetString(modulus,16)

	return fmt.Sprintf("%0256x", bigOrigData.Exp(bigOrigData, big.NewInt(exponent), bigModulus))

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
