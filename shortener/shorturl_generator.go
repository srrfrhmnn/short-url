package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortURL(initialLink string, userId string) string {
	// Hashing the initial link + the user id
	urlHashBytes := sha256Of(initialLink + userId)

	// Deriving a big integer number from the hash bytes
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Converting the hash bytes to a base58 encoded string and returning the first 8 characters
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	return finalString[:8]
}
