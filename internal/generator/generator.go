package generator

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func GenerateShortCode(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(
			rand.Reader, 
			big.NewInt(int64(len(charset))),
		)

		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}

	println(string(result))
	return string(result), nil
}
