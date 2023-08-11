package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func scrambleResponse(message string, publicKeyData []byte) (encryptedMessage string, err error) {
	// Encrypt plain text message using public key
	encryptedMessage, err = helper.EncryptMessageArmored(string(publicKeyData), message)
	if err != nil {
		return "", err
	}
	return encryptedMessage, nil
}

func deScrambledResponse(scrambledMessage string, passphrase []byte, privateKey string) (string, error) {
	// Decrypt armored encrypted message using the private key and obtain plain text
	decrypted, err := helper.DecryptMessageArmored(privateKey, passphrase, scrambledMessage)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}

func main() {
	// Replace with the correct paths to your PGP key files
	pgpPublicKeyFile := "path/to/publickey"
	pgpPrivateKeyFile := "path/to/privatekey"

	// Read the PGP public key from the file
	pgpPublicKeyData, err := ioutil.ReadFile(pgpPublicKeyFile)
	if err != nil {
		log.Fatal("Error reading PGP public key from file:", err)
	}

	// The message to be encrypted
	message := "Hello, World!"

	// Encrypt the message using the public key
	encryptedMsg, err := scrambleResponse(message, pgpPublicKeyData)
	if err != nil {
		log.Fatal("Error encrypting message:", err)
	}
	fmt.Println("Encrypted Message:", encryptedMsg)

	// Read the PGP private key from the file
	pgpPrivateKeyData, err := ioutil.ReadFile(pgpPrivateKeyFile)
	if err != nil {
		log.Fatal("Error reading PGP private key from file:", err)
	}

	// Decrypted the encrypted message using the private key
	decryptedMsg, err := deScrambledResponse(encryptedMsg, []byte("my_super_secret_passphrase"), string(pgpPrivateKeyData))
	if err != nil {
		log.Fatal("Error decrypting message:", err)
	}
	fmt.Println("Decrypted Message:", decryptedMsg)
}
