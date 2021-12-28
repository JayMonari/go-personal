package secret

import (
	"encoding/json"
	"errors"
	"os"
	"secret/cipher"
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
	if err := v.load(); err != nil {
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
	if err := v.load(); err != nil {
		return err
	}
	v.keyValues[key] = value
	if err := v.save(); err != nil {
		return err
	}
	return nil
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return json.NewDecoder(r).Decode(&v.keyValues)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := cipher.EncryptWriter(v.encodingKey, f)
	return json.NewEncoder(w).Encode(v.keyValues)
}
