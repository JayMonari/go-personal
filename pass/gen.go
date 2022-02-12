package pass

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const special = `~!@#$%^&*()_+-=\{[:;"'<,>.?/]}` + "`"
const digits = "0123456789"
const all = lower + upper + special + digits

// Gen generate a random password of given length n. The password includes
// lower, upper, special, and digit characters.
func Gen(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		bi, err := rand.Int(rand.Reader, big.NewInt(int64(len(all))))
		if err != nil {
			fmt.Printf("ERROR: %v", err)
			return ""
		}
		sb.WriteByte(all[bi.Int64()])
	}
	return sb.String()
}
