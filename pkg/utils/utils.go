package utils

import (
	"crypto/rand"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var buf [1]byte

func RandomID() string {
	s := make([]byte, 20)
	for i := 0; i < 20; i++ {
		_, _ = rand.Read(buf[:])
		s[i] = (chars[int(buf[0])%len(chars)])
	}
	return string(s)
}
