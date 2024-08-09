package utils

import (
	"math/rand"
	"time"
)

var SeededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

var charSet string = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[SeededRand.Intn(len(charSet))]
	}
	return string(b)
}
