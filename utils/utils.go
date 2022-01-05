package utils

import (
	"github.com/mr-tron/base58" // Invented with Bitcoin without character 0-zero O-oh l-el I- ii + /
	"log"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
