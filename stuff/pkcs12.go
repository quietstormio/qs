package stuff

import (
	"encoding/pem"
	"fmt"
	"os"

	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

func BigP12() {
	myFile, err := os.ReadFile("/path/to/cert/client-identity.p12")
	if err != nil {
		fmt.Println(err)
	}

	_, certificate, err := pkcs12.Decode(myFile, "Password")
	if err != nil {
		fmt.Println(err)
	}

	// Convert the public key to a string
	// If you want it to be a string
	// You can do stuff to the *x509.Certificate above like:
	// certificate.Subject
	publicKeyString := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: certificate.Raw,
	}))

	fmt.Print(publicKeyString)
}
