package crypto

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "errors"
    "math/big"
)

// Crypto defines the interface for cryptographic operations.
// This interface can be replaced with a PQC implementation
// or any other cryptographic scheme.
type Crypto interface {
    GenerateKeys() (privateKey interface{}, publicKey interface{}, err error)
    Sign(message []byte, privateKey interface{}) ([]byte, error)
    Verify(message []byte, signature []byte, publicKey interface{}) bool
}

// defaultCrypto uses ECDSA (secp256k1 or P256, etc.) as an example
type defaultCrypto struct{}

func NewDefaultCrypto() Crypto {
    return &defaultCrypto{}
}

func (dc *defaultCrypto) GenerateKeys() (interface{}, interface{}, error) {
    privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return nil, nil, err
    }
    return privKey, &privKey.PublicKey, nil
}

func (dc *defaultCrypto) Sign(message []byte, privateKey interface{}) ([]byte, error) {
    priv, ok := privateKey.(*ecdsa.PrivateKey)
    if !ok {
        return nil, errors.New("invalid private key type")
    }

    r, s, err := ecdsa.Sign(rand.Reader, priv, message)
    if err != nil {
        return nil, err
    }

    // Serialize r and s into a single byte slice for simplicity
    signature := append(r.Bytes(), s.Bytes()...)
    return signature, nil
}

func (dc *defaultCrypto) Verify(message []byte, signature []byte, publicKey interface{}) bool {
    pub, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return false
    }

    halfLen := len(signature) / 2
    r := new(big.Int).SetBytes(signature[:halfLen])
    s := new(big.Int).SetBytes(signature[halfLen:])

    return ecdsa.Verify(pub, message, r, s)
}
