package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int    // Same as Value in Output
	Sig string // Signature same as PubKey
}

func (tx *Transaction) SetID() {
	var (
		encoded bytes.Buffer
		hash    [32]byte
	)

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinBaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}
	txin := TxInput{[]byte{}, -1, data} // Reference no output -1
	txout := TxOutput{100, to}          // Rewards 100 Token

	tx := Transaction{
		ID:      nil,
		Inputs:  []TxInput{txin},
		Outputs: []TxOutput{txout},
	}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinBase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
