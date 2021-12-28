package secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"secret/encrypt"
	"strings"
	"sync"
)

type Vault struct {
	mu          sync.RWMutex
	encodingKey string
	filepath    string
	keyValues   map[string]string
}

func File(encodingKey, path string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    path,
	}
}

func (v *Vault) Get(key string) (string, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	if err := v.loadKeyValues(); err != nil {
		return "", err
	}
	if value, ok := v.keyValues[key]; ok {
		return value, nil
	}
	return "", errors.New("secret: no value for that key")
}

func (v *Vault) Set(key, value string) error {
	v.mu.Lock()
	defer v.mu.Unlock()
	if err := v.loadKeyValues(); err != nil {
		return err
	}
	v.keyValues[key] = value
	if err := v.saveKeyValues(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) loadKeyValues() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	var sb strings.Builder
	if _, err = io.Copy(&sb, f); err != nil {
		return err
	}
	decJSON, err := encrypt.Decrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}
	r := strings.NewReader(decJSON)
	if err = json.NewDecoder(r).Decode(&v.keyValues); err != nil {
		return err
	}
	return nil
}

func (v *Vault) saveKeyValues() error {
	var sb strings.Builder
	if err := json.NewEncoder(&sb).Encode(v.keyValues); err != nil {
		return err
	}

	encJSON, err := encrypt.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = fmt.Fprint(f, encJSON); err != nil {
		return err
	}
	return nil
}
