package link

import (
	"math/rand"
)

func generatePseudoLink() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, 6)
	for i := range key {
		key[i] = charset[rand.Intn(len(charset))]
	}
	return "https://short.ly/" + string(key) // Просто конкатенация
}
