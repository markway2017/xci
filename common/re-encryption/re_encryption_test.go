package re_encryption

import (
	"testing"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
)

func TestSet(t *testing.T) {

	size := 1024
	if testing.Short() {
		size = 128
	}
	patientPrivateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		t.Errorf("failed to generate key")
	}

	doctorPrivateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		t.Errorf("failed to generate key")
	}

	msg := "I am a patient"
	t.Log("Original data: "+msg)

	t.Log("Encrypt with patient public key")
	encryptedBytes,err := rsa.EncryptPKCS1v15(rand.Reader, &patientPrivateKey.PublicKey, []byte(msg))
	if err != nil {
		t.Errorf("Encrypt process error: "+err.Error())
	}
	t.Log(hex.EncodeToString(encryptedBytes))

	t.Log("Run re-encrypt")
	newEncryptedBytes,err := Reencrypt(encryptedBytes,patientPrivateKey,&doctorPrivateKey.PublicKey)
	if err != nil {
		t.Errorf("Reencrypt process error: "+err.Error())
	}
	t.Log(hex.EncodeToString(newEncryptedBytes))

	t.Log("Decrypt with new doctor private key")
	newMsg,err :=rsa.DecryptPKCS1v15(rand.Reader, doctorPrivateKey, newEncryptedBytes)
	if err != nil {
		t.Errorf("Decrypt process error: "+err.Error())
	}
	t.Log(string(newMsg))
}