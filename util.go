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
logics:

cases for on-chain
1. any valid address can write for single pre-defined topic. on-chain storage
2. any valid address can create a topic, and allow any other valid addresses to write to the topic. on-chain storage.
3. an allow-list to be added to case 2. The owner of each topic manages its allow-list, which contains a group of addresses.
*4. require allow-list addesses to register, and username, email, sex are required. (email can validated off-chain.) non-registered addesses will fail to write.
*5. each address can authorize another address to write. In this way, the second address is used only for the topic write (no token transfer for safety reason).
6. after APIs are done, the web UI is added. Since no wallet yet, the private key can be pasted directly at this moment. 

challenges:
1, what if it's a wallet?
2, always need private keys(need good key management)(mapping to password, better way) password storage and mapping
3, authorization methods. (?)
Summnary one:
challenges
(it does not make sense to have <username, email, sex>, we can allow self-defined ID which is either username or email. 
Then we will verify the <ID, priv_key> insteand of <address, priv_key>. It does not help much because the private key is still needed. 
What many users want should be username/password style.)
(priv_key vs password. for password, we can sign the data or decrypt the data)

cases for off-chain
1. use IPFS or IPFS-MFS
await ipfs.files.stat('/example')
This method returns an object with a cid, size, cumulativeSize, type, blocks, withLocality, local, and sizeLocal.
{
  hash: CID('QmXmJBmnYqXVuicUfn9uDCC8kxCEEzQpsAbeq1iJvLAmVs'),
  size: 60,
  cumulativeSize: 118,
  blocks: 1,
  type: 'directory'
}
2. JS needs to finish a transaction which contains update off-chain file and update smart contract index.
3. registration and allow-list logic should be same. Anyone can change the off-chain file, but it will not be accpeted unless it's in allow-list.
4. file based, or record based data storage.
5. (per file encryption, per record encryption need to be discussed below)


business user vs no business user to avoid censorship

$$$$$
1. if a project is using EVM smart contract, how to also use it?
EVM address should link to a COSMOS address (which can pay the gas fees if necessary). 
Then we can authorize the EVM address.

2. the concept of 'role'.

!!!!! if using EVM
EVM, cosmos, or EVMOS in cosmos
need to know if EVM can support all algorithm. If we use something other than hash functions.

3. Polygon

four categories: crypto user, non-crypto user, crypto developer, non-crypto developer

don't need to mention privacy too much.
maybe not focus on personal dada. otherwise zero knowledge is not enough

new design with new doc. review the new design and merge

MEETING NOTES:
1, SSO, ALEO
2, node management
3, key custodian.
4, customized ipfs
5, arbitrum code
6, less personal data management

TODO:
1, zero knowledge pass
2, key-value store, with encryption (single owner, multiple owners)
3, dynamic passcode

use Key Regression or symmetric link

trial:
chainlink oracle for data source, especially database.

FINAL:
zero knowledge + wallet ++ 
        evm + tendermint ++ 
                file + key-value ++ 
                        improved ipfs
------
user cases




============================================================================================================
cases for encryption
1. encrption: profile, message, authorizaion list?
2. 1 to 1, and 1 to many, many to many
3. authentication and identity

web-server for JS API hosting
mars
where to run? browser, wallet, webpage?
DNF

TODOs: a long list
duplicated storages
signature verification
user authentication
identity management

Ethereum compatible but inherit the limitation of Ethereum.
Incompatible and with more features.
Limitations: block speed, small storage, relying on smart contracts, no encryption in contracts.
Others: Binance Smart Chain. Avalanche
IPFS and Arweave
other storages

Reliability
Scalability
large scale
small scale
middle scale test

unit tests

layer0
layer1
layer2
layer3

readme in dot EM format
Here

client encryption or web-server 
others
no encryption mode
mixed encryption mode
combined encryption mode
traditional mode

design risk model
policy
enforcement function (related to the policies. need a mapping)
assessment of policy
compliance in different regions(7)
trust and safety
Risk
storage policy evaluation 
formula, factors
formulae versions
set
connection
massive scale deletions in the graph

zero-day issues, and the related windows.
code base, and it's builds, and apps that they depend on
graphs

use cases:
simple twitter 
java, rust, python(pyre, pysa), go, ruby, js
js web3, others

review
---------------------
both interactive and non-interactive authentication
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
/*openssl3.0.7 or backport*/
/*new supported algorithms and extra */
/*more algorithms */
/*FIPS 140-2 validation */
/*
the introduction of the Provider concept. Providers collect together and make available algorithm implementations. 
5 different providers as standard. 
One of the standard providers available is the FIPS provider. This makes available FIPS validated cryptographic algorithms.
Use of the low level APIs has been informally discouraged.
ENGINE API

Certificate Management Protocol (CMP, RFC 4210), CRMF (RFC 4211),  HTTP transfer (RFC 6712)
A proper HTTP(S) client in libcrypto supporting GET and POST, redirection, plain and ASN.1-encoded contents, proxies, and timeouts
*/
