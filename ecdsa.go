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

/* ECDSA signing*/
func EcdsaSign(privpem, input string) (rSig, sSig string, err error) {
        sum := sha256.Sum256([]byte(input))
        priv, err := parseEcdsaPrivateKeyFromPemStr(privpem)
        if (err != nil) {
                fmt.Println(err)
                return "", "", err
        }

        r, s, err := ecdsa.Sign(rand.Reader, priv, sum[:32])

        if err != nil {
                fmt.Println(err)
                return  "", "", err
        }

        return r.String(), s.String(), nil

}

