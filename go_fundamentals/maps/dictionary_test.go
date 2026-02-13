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
		t.Run("Update a Value", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			errUpt := dict.Update("test", "test passed")
			expected := "test passed"
			actual, errSch := dict.Search("test")

			require.NoError(errUpt)
			require.NoError(errSch)
			require.Equal(expected, actual)
		})
	})
	t.Run("Validation Errors", func(t *testing.T) {
		t.Run("Search Key Not Found", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{}

			_, err := dict.Search("test")

			require.Error(err)
			require.ErrorIs(err, ErrKeyNotFound)
		})
		t.Run("Add Void Key", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{}

			err := dict.Add("", "this is a test")

			require.Error(err)
			require.ErrorIs(err, ErrVoidKey)
		})
		t.Run("Add Key Already Exists", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			errAdd := dict.Add("test", "")

			require.Error(errAdd)
			require.ErrorIs(errAdd, ErrKeyAlreadyExists)
			require.Equal("this is a test", dict["test"])
		})
		t.Run("Update Void Key", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			errUpt := dict.Update("", "test did not pass")

			require.Error(errUpt)
			require.ErrorIs(errUpt, ErrVoidKey)
			require.Equal("this is a test", dict["test"])
		})
		t.Run("Update Key Not Found", func(t *testing.T) {
			require := require.New(t)
			dict := Dictionary{"test": "this is a test"}

			errUpt := dict.Update("not exists", "test did not pass")

			require.Error(errUpt)
			require.ErrorIs(errUpt, ErrKeyNotFound)
			require.Equal("this is a test", dict["test"])
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

func ExampleDictionary_Update() {
	dict := Dictionary{"test": "this is a test"}
	_ = dict.Update("test", "new value")
	fmt.Println(dict["test"])
	// Output: new value
}

func BenchmarkDictionary_Search(b *testing.B) {
	dict := Dictionary{"test": "this is a test"}
	for b.Loop() {
		_, _ = dict.Search("test")
	}
}

func BenchmarkDictionary_Add(b *testing.B) {
	for b.Loop() {
		dict := Dictionary{}
		_ = dict.Add("test", "this is a test")
	}
}

func BenchmarkDictionary_Update(b *testing.B) {
	dict := Dictionary{"test": "this is a test"}
	for b.Loop() {
		_ = dict.Update("test", "new value")
	}
}
