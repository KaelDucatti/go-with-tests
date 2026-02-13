package maps

const (
	ErrKeyNotFound      = DictionaryErr("key not found")
	ErrVoidKey          = DictionaryErr("key cannot be empty")
	ErrKeyAlreadyExists = DictionaryErr("key already exists")
)

type Dictionary map[string]string
type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key, value string) error {
	if key == "" {
		return ErrVoidKey
	}
	if _, exists := d[key]; exists {
		return ErrKeyAlreadyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if key == "" {
		return ErrVoidKey
	}
	if _, ok := d[key]; !ok {
		return ErrKeyNotFound
	}
	d[key] = value
	return nil
}
