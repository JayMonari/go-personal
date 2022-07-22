package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var currencies = [...]string{"EUR", "USD", "CAD"}

func init() { rand.Seed(time.Now().UnixNano()) }

// RandInt generates a random integer between min and max
func RandInt(min, max int64) int64 { return min + rand.Int63n(max-min+1) }

// RandString generate a random string of length n in lower ascii
func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		sb.WriteByte(alphabet[rand.Intn(k)])
	}
	return sb.String()
}

// RandOwner generates a random name for account owner
func RandOwner() string { return RandString(6) }

// RandBalance generates a random account balance
func RandBalance() int64 { return RandInt(0, 9001) }

func RandCurrency() string { return currencies[rand.Intn(len(currencies))] }
