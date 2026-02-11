package maps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDictionary(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("Known Key", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			expected := "this is a test"
			actual, err := dict.Search("test")

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("Add a New key-value", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{}

			errAdd := dict.Add("test", "this is a test")
			expected := "this is a test"
			actual, errSearch := dict.Search("test")

			require.NoError(errAdd)
			require.NoError(errSearch)
			require.Equal(expected, actual)
		})
	})
	t.Run("Validation Errors", func(t *testing.T) {
		t.Run("Unknown Key", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{}

			_, err := dict.Search("test")

			require.Error(err)
			require.EqualError(err, ErrKeyNotFound.Error())
		})
		t.Run("Add Void Key", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{}

			err := dict.Add("", "this is a test")

			require.Error(err)
			require.EqualError(err, ErrAddVoidKey.Error())
		})
		t.Run("Key Already Exists", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			err := dict.Add("test", "")

			require.Error(err)
			require.EqualError(err, ErrKeyAlreadyExists.Error())
		})
	})
}

func ExampleDictionary_Search() {
	dict := Dictionary{"test": "this is a test"}
	value, _ := dict.Search("test")
	fmt.Println(value)
	// Output: this is a test
}

func ExampleDictionary_Add() {
	dict := Dictionary{}
	_ = dict.Add("test", "this is a test")
	fmt.Println(dict["test"])
	// Output: this is a test
}

func BenchmarkDictionary_Search(b *testing.B) {
	for b.Loop() {
		dict := Dictionary{"test": "this is a test"}
		_, _ = dict.Search("test")
	}
}

func BenchmarkDictionary_Add(b *testing.B) {
	for b.Loop() {
		dict := Dictionary{}
		_ = dict.Add("test", "this is a test")
	}
}
