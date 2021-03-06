package server

import (
	"crypto/rand"
	"crypto/sha1"
)

func IdFromName(name string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(name))
	return hasher.Sum(nil)
}

func RandId() []byte {
	b := make([]byte, 16)
	rand.Read(b)
	return b
}
