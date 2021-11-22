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

func RsaVerify(pubpem, input, signature string) bool{
        var opts rsa.PSSOptions
        opts.SaltLength = rsa.PSSSaltLengthAuto
        rsaPublicKey, err := parseRsaPublicKeyFromPemStr(pubpem)
        if (err != nil) {
                fmt.Println(err)
                return false
        }

        sum := sha256.Sum256([]byte(input))
        err = rsa.VerifyPSS(rsaPublicKey, crypto.SHA256, sum[:32], []byte(signature), &opts)

        if err != nil {
                fmt.Println(err)
                return false
        }

        return true
}

