package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"log"

	"golang.org/x/crypto/ssh"
)

func generatePublicKeyAuthMethod() (ssh.AuthMethod, error) {
	var _, privKey, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	privBlock, err := ssh.MarshalPrivateKey(crypto.PrivateKey(privKey), "")
	if err != nil {
		return nil, err
	}

	privatePEM := pem.EncodeToMemory(privBlock)

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
