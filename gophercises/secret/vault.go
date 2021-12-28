package secret

import (
	"errors"
	"secret/encrypt"
)

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func New(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues:   make(map[string]string),
	}
}

func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	val, err := encrypt.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (v *Vault) Set(key, value string) error {
	encVal, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	v.keyValues[key] = encVal
	return nil
}
