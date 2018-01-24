package sphinx

import (
	"crypto/sha256"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/elliptic"
	"crypto/rand"
)

func AES_CTR(key, plaintext []byte) []byte {

	ciphertext := make([]byte, len(plaintext))

	iv := []byte("0000000000000000")
	//if _, err := io.ReadFull(crand.Reader, iv); err != nil {
	//	panic(err)
	//}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func hash(arg []byte) []byte{

	h := sha256.New()
	h.Write(arg)

	return h.Sum(nil)
}

func Hmac(key, message []byte) []byte{
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func GenerateKeyPair() ([]byte, []byte){
	priv, x, y, err  := elliptic.GenerateKey(elliptic.P224(), rand.Reader)

	if err != nil {
		panic(err)
	}

	return elliptic.Marshal(elliptic.P224(), x, y), priv
}