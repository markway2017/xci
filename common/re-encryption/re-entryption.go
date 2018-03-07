package re_encryption

import (
	"crypto/rsa"
	"crypto/rand"
	"errors"
)

func Reencrypt(data []byte, patientPrivate *rsa.PrivateKey, doctorPublic *rsa.PublicKey) ([]byte, error) {

	decryptedBytes,err :=rsa.DecryptPKCS1v15(rand.Reader, patientPrivate, data)
	if err != nil {
		return nil, errors.New("Decrypt with patient private key error: "+err.Error())
	}

	encryptedBytes,err := rsa.EncryptPKCS1v15(rand.Reader, doctorPublic, decryptedBytes)
	if err != nil {
		return nil, errors.New("Encrypt with doctor public key error: "+err.Error())
	}

	return encryptedBytes,nil
}