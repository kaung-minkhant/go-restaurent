package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot hash the password")
	}
	return string(encrypted), nil
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func SignData(data []byte, privatePath *string) (string, error) {
	privateKey, err := parseRsaPrivateKeyFromPemFile(privatePath)
	if err != nil {
		return "", err
	}
	msgHash := sha256.New()
	_, err = msgHash.Write(data)
	if err != nil {
		return "", err
	}
	dataHashSum := msgHash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dataHashSum)
	if err != nil {
		return "", err
	}
	encoded := hex.EncodeToString(signature)
	return encoded, nil
}

func VerifySign(data []byte, signature string, pubPath *string) error {
	publicKey, err := parseRsaPublicKeyFromPemFile(pubPath)
	if err != nil {
		return err
	}
	dataHash := sha256.New()
	_, err = dataHash.Write(data)
	if err != nil {
		panic(err)
	}
	dataHashSum := dataHash.Sum(nil)
	decodedSig, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, dataHashSum, decodedSig)
}

func parseRsaPrivateKeyFromPemFile(path *string) (*rsa.PrivateKey, error) {
	var privateKeyPath string
	if path == nil {
		privateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
	} else {
		privateKeyPath = *path
	}
	file, err := os.Open(privateKeyPath)
	if err != nil {
		return nil, err
	}
	fileContent := make([]byte, 0)
	for {
		fileChunk := make([]byte, 512)
		n, err := file.Read(fileChunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		fileContent = append(fileContent, fileChunk[:n]...)
	}
	block, _ := pem.Decode(fileContent)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func parseRsaPublicKeyFromPemFile(path *string) (*rsa.PublicKey, error) {
	var publicKeyPath string
	if path == nil {
		publicKeyPath = os.Getenv("PUBLIC_KEY_PATH")
	} else {
		publicKeyPath = *path
	}
	file, err := os.Open(publicKeyPath)
	if err != nil {
		return nil, err
	}
	fileContent := make([]byte, 0)
	for {
		fileChunk := make([]byte, 512)
		n, err := file.Read(fileChunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		fileContent = append(fileContent, fileChunk[:n]...)
	}
	block, _ := pem.Decode(fileContent)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}
