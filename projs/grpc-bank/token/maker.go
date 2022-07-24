package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken makes a new token for username for certain duration.
	CreateToken(username string, d time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not.
	VerifyToken(token string) (*Payload, error)
}
