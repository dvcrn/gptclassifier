package utils

import (
	"log"

	tokenizer "github.com/sandwich-go/gpt3-encoder"
)

func CountTokens(input string) int {
	encoder, err := tokenizer.NewEncoder()
	if err != nil {
		log.Fatal(err)
	}

	encoded, err := encoder.Encode(input)
	if err != nil {
		log.Fatal(err)
	}

	return len(encoded)
}

func SliceTokens(input string, tokenCount int) string {
	encoder, err := tokenizer.NewEncoder()
	if err != nil {
		log.Fatal(err)
	}

	encoded, err := encoder.Encode(input)
	if err != nil {
		log.Fatal(err)
	}

	if tokenCount >= len(encoded) {
		return input
	}

	decoded := encoder.Decode(encoded[:tokenCount-1])
	return decoded
}
