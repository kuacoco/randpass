package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func main() {

	var (
		num = flag.Int(
			"n", 8, "Password Length",
		)
	)
	flag.Parse()

	fmt.Printf("%s\n", randStringBytes(*num))
}

func randStringBytes(n int) []byte {

	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return b
}
