package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/newbiet21379/blockgo/utils"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const (
	checkSumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w *Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionHash := append([]byte{version}, pubHash...)
	checksum := CheckSum(versionHash)

	fullHash := append(versionHash, checksum...)
	address := utils.Base58Encode(fullHash)

	ValidateAddress(fmt.Sprintf("%s", address))

	return address
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	Wallet := Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}
	return &Wallet
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func CheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checkSumLength]
}

func ValidateAddress(address string) bool {
	pubKeyHash := utils.Base58Decode([]byte(address))
	actualCheckSum := pubKeyHash[len(pubKeyHash)-checkSumLength:]

	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checkSumLength]
	targetCheckSum := CheckSum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualCheckSum, targetCheckSum) == 0
}
