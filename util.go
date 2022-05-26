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

/*
cases
1. any valid address can write for single pre-defined topic. on-chain storage
2. any valid address can create a topic, and allow any other valid addresses to write to the topic. on-chain storage.
3. an allow-list to be added to case 2. The owner of each topoc manages its allow-list, which contains a group of addresses.
4. require allow-list addesses to register, and username, email, sex are required. (email can validated off-chain. non-registered addesses will fail to write.
5. 
6. after APIs are done, the web UI is added. Since no wallet yet, the private key can be pasted directly at this moment. 
)

*/

/* util functions */
func parseAll()(*rsa.PrivateKey, error){
        return nil, nil
}

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

/*AES256-GCM
GCM is safer then CBC, and well adopted.
Plus, GCM allows associated additional data(AAD).
*/

/*Chacha-Poly1035*/
