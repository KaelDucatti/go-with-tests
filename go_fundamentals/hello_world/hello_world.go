package hello_world

import "fmt"

var greetings = map[string]string {
	"es"	: "Hola",
	"pt-br"	: "Olá", 
	"en"	: "Hello",
}

func GreetingByLang(lang string) (value string) {
	if value, ok := greetings[lang]; ok {
		return value
	}
	return "Hello"
}

func Greeting(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s", GreetingByLang(lang), name)

}