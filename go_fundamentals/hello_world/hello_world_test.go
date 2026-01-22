package hello_world

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
		assert := assert.New(t)

		got := Greeting("Kael", "")
		want := "Hello, Kael"
		
		assert.Equal(want, got)
	})

	t.Run("saying hello without name", func(t *testing.T){
		assert := assert.New(t)

		got := Greeting("", "pt-br")
		want := "Olá, World"

		assert.Equal(want, got)
	})
	
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		assert := assert.New(t)

		got := Greeting("", "")
		want := "Hello, World"

		assert.Equal(want, got)
	})

	t.Run("in Spanish", func(t *testing.T){
		assert := assert.New(t)

		got := Greeting("Kael", "es")
		want := "Hola, Kael"

		assert.Equal(want, got)
	})

	t.Run("in Portuguese", func(t *testing.T){
		assert := assert.New(t)

		got := Greeting("Kael", "pt-br")
		want := "Olá, Kael"
		
		assert.Equal(want, got)
	})
}

func ExampleGreeting() {
	greet := Greeting("Kael", "pt-br")
	fmt.Println(greet)
	// Output: Olá, Kael
}