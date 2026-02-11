package maps

import "errors"

var (
	ErrWordNotFound    error = errors.New("key not found")
	ErrAddVoidKey      error = errors.New("key value cannot be void")
	ErrKeyAlreadyExits error = errors.New("key already exists in dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrWordNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	if key == "" {
		return ErrAddVoidKey
	}
	if _, exists := d[key]; exists {
		return ErrKeyAlreadyExits
	}
	d[key] = value
	return nil
}
