package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"os"
)

type KeyGenConfig struct {
	path               string // path relative to this file or absolute path
	privateKeyFilename string
	publicKeyFilename  string
}

var Config = &KeyGenConfig{
	path:               "../keys",
	privateKeyFilename: "privKey.pem",
	publicKeyFilename:  "pubKey.pem",
}

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ExportRsaPrivateKeyAsPemFile(privkey *rsa.PrivateKey) error {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	filePath := Config.path + "/" + Config.privateKeyFilename
	privFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer privFile.Close()
	return pem.Encode(
		privFile,
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParseRsaPrivateKeyFromPemFile() (*rsa.PrivateKey, error) {
	file, err := os.Open("../keys/privKey.pem")
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

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ExportRsaPublicKeyAsPemFile(pubkey *rsa.PublicKey) error {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return err
	}
	filePath := Config.path + "/" + Config.publicKeyFilename
	pubFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer pubFile.Close()
	return pem.Encode(
		pubFile,
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
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

func ParseRsaPublicKeyFromPemFile() (*rsa.PublicKey, error) {
	file, err := os.Open("../keys/pubKey.pem")
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

func main() {

	// // // Create the keys
	// priv, pub := GenerateRsaKeyPair()

	// // // Export the keys to pem string
	// if err := ExportRsaPrivateKeyAsPemFile(priv); err != nil {
	// 	log.Fatal("Exporting private key failed:", err)
	// }
	// if err := ExportRsaPublicKeyAsPemFile(pub); err != nil {
	// 	log.Fatal("Exporting public key failed:", err)
	// }
	// privFromFile, err := ParseRsaPrivateKeyFromPemFile()
	// if err != nil {
	// 	fmt.Println("private from file error", err)
	// }
	// pubFromFile, err := ParseRsaPublicKeyFromPemFile()
	// if err != nil {
	// 	fmt.Println("public from file error", err)
	// }
	// // fmt.Println("private", priv)
	// fmt.Println("private from file", privFromFile)
	// fmt.Println("public from file", pubFromFile)
}
