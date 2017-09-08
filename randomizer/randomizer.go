package randomizer

import (
	"math/rand"
	"time"
)

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	r := make([]rune, n)
	for i := range r {
		r[i] = runes[rand.Intn(len(runes))]
	}
	return string(r)
}
