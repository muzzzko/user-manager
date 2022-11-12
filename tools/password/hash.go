package password

import (
	"crypto/sha512"
	"fmt"
)

func Hash(password string) (string, error) {
	hash := sha512.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("fail create hash: %w", err)
	}
	passwordHash := fmt.Sprintf("%x", hash.Sum(nil))

	return passwordHash, nil
}
