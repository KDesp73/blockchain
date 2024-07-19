package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

func Hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func SignData(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func FormatKeys(public *rsa.PublicKey, private *rsa.PrivateKey) (string, string, error) {
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(private)})
	publicKeyBytes, err := ssh.NewPublicKey(public)
	if err != nil {
		return "", "", err
	}
	publicKeyPEM := ssh.MarshalAuthorizedKey(publicKeyBytes)

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

func FormatPublicKey(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	return string(ssh.MarshalAuthorizedKey(publicKeyBytes)), nil
}

func GenerateRSAKeys() (*rsa.PublicKey, *rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	

	return &privateKey.PublicKey, privateKey, nil
}

func VerifySignature(publicKey *rsa.PublicKey, data []byte, signature []byte) error {
	hashed := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return err
	}
	return nil
}
