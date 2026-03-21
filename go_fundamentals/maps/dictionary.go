package maps

const (
	ErrKeyNotFound      = DictionaryErr("Key Not Found in Dictionary")
	ErrKeyAlreadyExists = DictionaryErr("Key already exists in dictionary")
)

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

type Dictionary map[string]string

func NewDictionary() Dictionary {
	return make(Dictionary)
}

func (d Dictionary) Search(key string) (string, error) {
	value, exists := d[key]
	if !exists {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, exists := d[key]
	if exists {
		return ErrKeyAlreadyExists
	}
	d[key] = value
	return nil
}