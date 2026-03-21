package maps_test

import (
	"fmt"
	"testing"

	. "github.com/KaelDucatti/go-with-tests/go_fundamentals/maps"
	"github.com/stretchr/testify/require"
)


func TestDictionarySuccessCases(t *testing.T) {
	t.Run("should return the value of the key", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{"test": "this is a test"}
		
		expected := "this is a test"
		actual, err := dict.Search("test")

		require.NoError(err)
		require.Equal(expected, actual)
	})
	t.Run("should add new key-value into the dictionary", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{}

		errAdd := dict.Add("test", "this is a test")
		expected := "this is a test"
		actual, errSearch := dict.Search("test")

		require.NoError(errAdd)
		require.NoError(errSearch)
		require.Equal(expected, actual)
	})
	t.Run("should update the key value of the dictionary", func(t *testing.T) {
		require := require.New(t)
		key := "test"
		value := "this is a test"
		dict := Dictionary{key: value}

		errUpdate := dict.Update(key, "still a test")
		expected := "still a test"
		actual, errSearch := dict.Search(key)

		require.NoError(errUpdate)
		require.NoError(errSearch)
		require.Equal(expected, actual)
	})
	t.Run("should delete the key-value of the dictionary", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{"test": "this is a test"}

		errDelete := dict.Delete("test")
		_, errSearch := dict.Search("test")

		require.NoError(errDelete)
		require.Error(errSearch)
		require.ErrorIs(errSearch, ErrKeyNotFoundOnSearch)
	})
}

func TestDictionaryValidationErrors(t *testing.T) {
	t.Run("should return error when key does not exist (search)", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{"test": "this is a test"}

		_, err := dict.Search("unknown")

		require.Error(err)
		require.ErrorIs(err, ErrKeyNotFoundOnSearch)
	})
	t.Run("should return error when key already exist", func(t *testing.T) {
		require := require.New(t)	
		key := "test"
		value := "this is a test"
		dict := Dictionary{key: value}

		errAdd := dict.Add(key, "still a test")
		postValue, errSearch := dict.Search("test")

		require.Error(errAdd)
		require.NoError(errSearch)
		require.ErrorIs(errAdd, ErrKeyAlreadyExists)
		require.Equal(value, postValue)
	})
	t.Run("should return error when key does not exist (update)", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{"test": "this is a test"}

		err := dict.Update("unknown", "still a test")

		require.Error(err)
		require.ErrorIs(err, ErrKeyNotFoundOnUpdate)
	})
	t.Run("should return error when key does not exist (delete)", func(t *testing.T) {
		require := require.New(t)
		dict := Dictionary{"test": "this is a test"}

		err := dict.Delete("unknown")

		require.Error(err)
		require.ErrorIs(err, ErrKeyNotFoundOnDelete)
	})
}

func ExampleNewDictionary() {
	dict := NewDictionary()
	fmt.Println(dict)
	// Output: map[]
}

func ExampleDictionary_Add() {
	dict := NewDictionary()
	dict.Add("test", "this is a test")
	fmt.Println(dict)
	// Output: map[test:this is a test]
}

func ExampleDictionary_Search() {
	dict := Dictionary{"test": "this is a test"}
	value, _ := dict.Search("test")
	fmt.Println(value)
	// Output: this is a test
}

func ExampleDictionary_Update() {
	dict := Dictionary{"test": "this is a test"}
	dict.Update("test", "still a test")
	fmt.Println(dict)
	// Output: map[test:still a test]
}

func ExampleDictionary_Delete() {
	dict := Dictionary{"test": "this is a test"}
	dict.Delete("test")
	fmt.Println(dict)
	// Output: map[]
}

func BenchmarkNewDictionary(b *testing.B) {
	for b.Loop() {
		_ = NewDictionary()
	}
}

func BenchmarkDictionary_Add(b *testing.B) {
	for b.Loop() {
		dict := NewDictionary()
		dict.Add("test", "this is a test")
	}
}

func BenchmarkDictionary_Search(b *testing.B) {
	dict := Dictionary{"test": "this is a test"}
	for b.Loop() {
		_, _ =dict.Search("test")
	}
}

func BenchmarkDictionary_Update(b *testing.B) {
	dict := Dictionary{"test": "this is a test"}
	for b.Loop() {
		dict.Update("test", "still a test")
	}
}

func BenchmarkDictionary_Delete(b *testing.B) {
	for b.Loop() {
		dict := Dictionary{"test": "this is a test"}
		dict.Delete("test")
	}
}