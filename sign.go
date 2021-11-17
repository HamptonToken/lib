package sign

import (
	"crypto"
	"crypto/x509"
	"crypto/rsa"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"errors"
	"crypto/sha256"
	"crypto/rand"
	"encoding/pem"
)


func RsaSign(privpem, input string) (string, error) {
	sum := sha256.Sum256([]byte(input))
	priv, err := parseRsaPrivateKeyFromPemStr(privpem)
	if (err != nil) {
		fmt.Println(err)
		return "", nil
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, sum[:32], &opts)

	if err != nil {
		fmt.Println(err)
		return  "", err
	}

	return string(signature), nil

}
