package re_encryption

import (
	"crypto/rand"
	"errors"
	"github.com/xcareteam/xci/crypto/ecies"
	"crypto/ecdsa"
)

func Reencrypt(data []byte, patientPrivate *ecdsa.PrivateKey, doctorPublic *ecdsa.PublicKey) ([]byte, error) {

	patientEciesPrivate := ecies.ImportECDSA(patientPrivate)

	doctorEciesPublic := ecies.ImportECDSAPublic(doctorPublic)

	decryptedBytes,err :=patientEciesPrivate.Decrypt(rand.Reader, data, nil,nil)
	if err != nil {
		return nil, errors.New("Decrypt with patient private key error: "+err.Error())
	}

	encryptedBytes,err := ecies.Encrypt(rand.Reader, doctorEciesPublic, decryptedBytes, nil, nil)
	if err != nil {
		return nil, errors.New("Encrypt with doctor public key error: "+err.Error())
	}

	return encryptedBytes,nil
}