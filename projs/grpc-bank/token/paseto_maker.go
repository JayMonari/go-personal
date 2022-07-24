package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is a PASETO token maker.
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker.
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters",
			chacha20poly1305.KeySize)
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

// CreateToken implements Maker
func (m PasetoMaker) CreateToken(username string, d time.Duration) (string, error) {
	p, err := NewPayload(username, d)
	if err != nil {
		return "", err
	}
	return m.paseto.Encrypt(m.symmetricKey, p, nil)
}

// VerifyToken implements Maker
func (m PasetoMaker) VerifyToken(token string) (*Payload, error) {
	p := &Payload{}
	if err := m.paseto.Decrypt(token, m.symmetricKey, p, nil); err != nil {
		return nil, ErrInvalidToken
	}
	if err := p.Valid(); err != nil {
		return nil, err
	}
	return p, nil
}
