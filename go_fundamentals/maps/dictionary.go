package maps

const (
	ErrKeyNotFoundOnSearch = DictionaryErr("Key Not Found in Dictionary")
	ErrKeyAlreadyExists    = DictionaryErr("Key already exists in dictionary")
	ErrKeyNotFoundOnUpdate = DictionaryErr("Cannot update the key's value because it doesn't exist")
	ErrKeyNotFoundOnDelete = DictionaryErr("Cannot delete a key that does not exist")
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
		return "", ErrKeyNotFoundOnSearch
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFoundOnSearch:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFoundOnSearch:
		return ErrKeyNotFoundOnUpdate
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotFoundOnSearch:
		return ErrKeyNotFoundOnDelete
	case nil:
		delete(d, key)
	default:
		return err
	}
	return nil
}