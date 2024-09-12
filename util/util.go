package util

import (
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	r = rand.New(src)
}

func RandomInt(min int64, max int64) int64 {
	return min + int64(r.Int63()%(int64(max-min+1)))
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)
	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(9)
}

func RandomPassword() string {
	return RandomString(12)
}
