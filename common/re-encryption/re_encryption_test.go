package re_encryption

import (
	"testing"
	"crypto/rand"
	"encoding/hex"
	"crypto/elliptic"
	"crypto/ecdsa"
	"github.com/xcareteam/xci/crypto/ecies"
	"time"
)

func TestReencrypt(t *testing.T) {

	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256

	patientPrivateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	if err != nil {
		t.Errorf("failed to generate key")
		t.Fail()
	}

	doctorPrivateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	if err != nil {
		t.Errorf("failed to generate key")
		t.Fail()
	}

	start := time.Now()

	msg := "I am a patient"
	t.Log("Original data: "+msg)

	t.Log("Encrypt with patient public key")
	patientEciesPrivate := ecies.ImportECDSA(patientPrivateKey)
	encryptedBytes,err := ecies.Encrypt(rand.Reader, &patientEciesPrivate.PublicKey, []byte(msg), nil, nil)
	if err != nil {
		t.Errorf("Encrypt process error: "+err.Error())
		t.Fail()
	}
	t.Log(hex.EncodeToString(encryptedBytes))

	t.Log("Run re-encrypt")
	newEncryptedBytes,err := Reencrypt(encryptedBytes,patientPrivateKey,&doctorPrivateKey.PublicKey)
	if err != nil {
		t.Errorf("Reencrypt process error: "+err.Error())
		t.Fail()
	}
	t.Log(hex.EncodeToString(newEncryptedBytes))

	t.Log("Decrypt with doctor private key")
	doctorEciesPrivate := ecies.ImportECDSA(doctorPrivateKey)
	newMsg,err :=doctorEciesPrivate.Decrypt(rand.Reader, newEncryptedBytes,nil,nil)
	if err != nil {
		t.Errorf("Decrypt process error: "+err.Error())
		t.Fail()
	}
	t.Log(string(newMsg))

	if string(newMsg) != msg {
		t.Fail()
	}

	elapsed := time.Since(start)
	t.Logf("Time cost %s", elapsed)
}