package auth

import (
	"encoding/json"
	"errors"

	"github.com/zalando/go-keyring"
)

const (
	keyAPIKey       = "api_key"
	keyTokenInfo    = "token_info"
	keyClientID     = "client_id"
	keyClientSecret = "client_secret"
)

// Storage defines the interface for credential storage
type Storage interface {
	// API Key methods
	GetAPIKey() (string, error)
	SetAPIKey(key string) error
	DeleteAPIKey() error

	// OAuth token methods
	GetTokenInfo() (*TokenInfo, error)
	SetTokenInfo(info *TokenInfo) error
	DeleteTokenInfo() error

	// Client credentials methods
	GetClientID() (string, error)
	SetClientID(id string) error
	DeleteClientID() error

	GetClientSecret() (string, error)
	SetClientSecret(secret string) error
	DeleteClientSecret() error
}

// KeyringStorage implements Storage using the system keyring
type KeyringStorage struct {
	service string
}

// NewKeyringStorage creates a new keyring-based storage
func NewKeyringStorage() *KeyringStorage {
	return &KeyringStorage{
		service: ServiceName,
	}
}

// GetAPIKey retrieves the stored API key
func (s *KeyringStorage) GetAPIKey() (string, error) {
	return keyring.Get(s.service, keyAPIKey)
}

// SetAPIKey stores an API key
func (s *KeyringStorage) SetAPIKey(key string) error {
	return keyring.Set(s.service, keyAPIKey, key)
}

// DeleteAPIKey removes the stored API key
func (s *KeyringStorage) DeleteAPIKey() error {
	err := keyring.Delete(s.service, keyAPIKey)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// GetTokenInfo retrieves stored OAuth token info
func (s *KeyringStorage) GetTokenInfo() (*TokenInfo, error) {
	data, err := keyring.Get(s.service, keyTokenInfo)
	if err != nil {
		return nil, err
	}

	var info TokenInfo
	if err := json.Unmarshal([]byte(data), &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// SetTokenInfo stores OAuth token info
func (s *KeyringStorage) SetTokenInfo(info *TokenInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return keyring.Set(s.service, keyTokenInfo, string(data))
}

// DeleteTokenInfo removes stored OAuth token info
func (s *KeyringStorage) DeleteTokenInfo() error {
	err := keyring.Delete(s.service, keyTokenInfo)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// GetClientID retrieves the stored client ID
func (s *KeyringStorage) GetClientID() (string, error) {
	return keyring.Get(s.service, keyClientID)
}

// SetClientID stores a client ID
func (s *KeyringStorage) SetClientID(id string) error {
	return keyring.Set(s.service, keyClientID, id)
}

// DeleteClientID removes the stored client ID
func (s *KeyringStorage) DeleteClientID() error {
	err := keyring.Delete(s.service, keyClientID)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// GetClientSecret retrieves the stored client secret
func (s *KeyringStorage) GetClientSecret() (string, error) {
	return keyring.Get(s.service, keyClientSecret)
}

// SetClientSecret stores a client secret
func (s *KeyringStorage) SetClientSecret(secret string) error {
	return keyring.Set(s.service, keyClientSecret, secret)
}

// DeleteClientSecret removes the stored client secret
func (s *KeyringStorage) DeleteClientSecret() error {
	err := keyring.Delete(s.service, keyClientSecret)
	if errors.Is(err, keyring.ErrNotFound) {
		return nil
	}
	return err
}

// MemoryStorage implements Storage using in-memory maps (for testing)
type MemoryStorage struct {
	data map[string]string
}

// NewMemoryStorage creates a new in-memory storage (for testing)
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (s *MemoryStorage) GetAPIKey() (string, error) {
	if v, ok := s.data[keyAPIKey]; ok {
		return v, nil
	}
	return "", keyring.ErrNotFound
}

func (s *MemoryStorage) SetAPIKey(key string) error {
	s.data[keyAPIKey] = key
	return nil
}

func (s *MemoryStorage) DeleteAPIKey() error {
	delete(s.data, keyAPIKey)
	return nil
}

func (s *MemoryStorage) GetTokenInfo() (*TokenInfo, error) {
	data, ok := s.data[keyTokenInfo]
	if !ok {
		return nil, keyring.ErrNotFound
	}
	var info TokenInfo
	if err := json.Unmarshal([]byte(data), &info); err != nil {
		return nil, err
	}
	return &info, nil
}

func (s *MemoryStorage) SetTokenInfo(info *TokenInfo) error {
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	s.data[keyTokenInfo] = string(data)
	return nil
}

func (s *MemoryStorage) DeleteTokenInfo() error {
	delete(s.data, keyTokenInfo)
	return nil
}

func (s *MemoryStorage) GetClientID() (string, error) {
	if v, ok := s.data[keyClientID]; ok {
		return v, nil
	}
	return "", keyring.ErrNotFound
}

func (s *MemoryStorage) SetClientID(id string) error {
	s.data[keyClientID] = id
	return nil
}

func (s *MemoryStorage) DeleteClientID() error {
	delete(s.data, keyClientID)
	return nil
}

func (s *MemoryStorage) GetClientSecret() (string, error) {
	if v, ok := s.data[keyClientSecret]; ok {
		return v, nil
	}
	return "", keyring.ErrNotFound
}

func (s *MemoryStorage) SetClientSecret(secret string) error {
	s.data[keyClientSecret] = secret
	return nil
}

func (s *MemoryStorage) DeleteClientSecret() error {
	delete(s.data, keyClientSecret)
	return nil
}
