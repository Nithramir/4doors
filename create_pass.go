package main

import (
	"crypto/rand"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) []byte {
	b := make([]byte, n)
	n, err := rand.Read(b)
	if err != nil {
		handle_err(err)
	}
	return b
}
