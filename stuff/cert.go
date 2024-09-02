package stuff

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func GetPasswordEnv() string {
	userTokenPassword := os.Getenv("QSPASSWORD")
	if !(len(userTokenPassword) > 0) {
		log.Fatal("Could not find your password ENV")
	}
	return (userTokenPassword)
}

func ReadCert(myPublicKeyPath string, myPrivateKeyPath string) {
	//how to get first character in string
	fmt.Println(myPublicKeyPath[0:1])
	fmt.Println(myPrivateKeyPath[0:1])

	//get home directory of current user
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	//read file, don't need to close (as far as we know)
	public, err := os.ReadFile(homeDirectory + "/Documents/certs/publickey.crt")
	if err != nil {
		fmt.Println("Error", err)
	}

	private, err := os.ReadFile(homeDirectory + "/Documents/certs/pkcs8.key")
	if err != nil {
		fmt.Println("Error", err)
	}

	//decode the string, can't just pass string from file
	spkiBlock, _ := pem.Decode([]byte(string(public)))
	privateBlock, _ := pem.Decode([]byte(string(private)))

	var spkiKey *rsa.PublicKey
	var spkiPrivate *rsa.PrivateKey

	pubInterface, _ := x509.ParsePKIXPublicKey(spkiBlock.Bytes)
	privateKey, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)

	if err != nil {
		panic(err)
	}
	spkiKey = pubInterface.(*rsa.PublicKey)
	spkiPrivate = privateKey.(*rsa.PrivateKey)

	plainText := []byte("This is a secret message")
	oaepLabel := []byte("")
	oaepDigests := sha256.New()

	cipherText, _ := rsa.EncryptOAEP(oaepDigests, rand.Reader, spkiKey, plainText, oaepLabel)
	decryptedValue, err := rsa.DecryptOAEP(oaepDigests, rand.Reader, spkiPrivate, cipherText, oaepLabel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decryptedValue)
}
