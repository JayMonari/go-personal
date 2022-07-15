package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "0123456789"
const special = `!@#$%^&*()_+-=[]{}\|;:'"<,>./?` + "`"
const passLen = 16

func main() {
	charSet := lower + upper + digits + special
	var sb strings.Builder
	for i := 0; i < passLen; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			log.Fatalf("Could not get random bigInt: %v", err)
		}
		sb.WriteByte(charSet[n.Int64()])
	}
	fmt.Println(sb.String())
}
