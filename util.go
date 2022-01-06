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

/* util functions */

/* parse the PKCS format private key */
func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the privkey")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

/* parse private key*/
func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the privkey")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}
/* parse public key*/
func parseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
        block, _ := pem.Decode([]byte(pubPEM))
        if block == nil {
                fmt.Println("failed to parse certificate PEM")
                return nil, errors.New("failed to parse PEM block containing the cert")
        }
        cert, err := x509.ParseCertificate(block.Bytes)
        if err != nil {
                fmt.Println("failed to parse certificate: ", err.Error())
                return nil, err
        }

        pub := cert.PublicKey.(*rsa.PublicKey)

        return pub, nil
}

/* others */

/*AES256-GCM*/

/*Chacha-Poly1035*/
