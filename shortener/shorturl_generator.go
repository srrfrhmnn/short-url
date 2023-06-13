package shortener

import (
	"crypto/sha256"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}
