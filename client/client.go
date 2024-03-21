package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"

	"golang.org/x/crypto/ssh"
)

func generatePublicKeyAuthMethod() (ssh.AuthMethod, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	privatePEM := pem.EncodeToMemory(&privBlock)

	signer, err := ssh.ParsePrivateKey(privatePEM)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(signer), nil
}

func main() {
	authMethod, err := generatePublicKeyAuthMethod()
	if err != nil {
		log.Fatalf("Failed to generate public key: %s", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	_, err = ssh.Dial("tcp", "127.0.0.1:8022", sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
	}

	log.Println("Connected to SSH server")
}
