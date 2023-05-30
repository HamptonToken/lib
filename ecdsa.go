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

/* ECDSA signing the input string, sha256*/
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

/* ECDSA verifying the input string with sig1 and sig2*/
func EcdsaVerify(pubpem, input, signature1, signature2 string) bool{

        rSig := new(big.Int)
        rSig, ok := rSig.SetString(signature1, 10)
        if !ok {
                fmt.Println("SetString: error")
                return false
        }

        sSig := new(big.Int)
        sSig, ok = sSig.SetString(signature2, 10)
        if !ok {
                fmt.Println("SetString: error")
                return false
        }

        ecPublicKey, err := parseEcdsaPublicKeyFromPemStr(pubpem)
        if (err != nil) {
                fmt.Println(err)
                return false
        }

        sum := sha256.Sum256([]byte(input))
        valid := ecdsa.Verify(ecPublicKey, sum[:32], rSig, sSig)

        return valid
}

/* Edwards-curve Digital Signature Algorithm (EdDSA) is also needed here*/
/* */
/* Ed25519 is the EdDSA signature scheme using SHA-512 (SHA-2) and Curve25519 */
