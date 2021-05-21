package hello_test

import (
	"fmt"

	"github.com/alemelomeza/go-cli/pkg/hello"
)

func ExampleHello_empty() {
	fmt.Println(hello.Hello("", ""))
	// Output: Hello, World
}

func ExampleHello_english() {
	fmt.Println(hello.Hello("en", "John"))
	// Output: Hello, John
}

func ExampleHello_spanish() {
	fmt.Println(hello.Hello("es", "Juan"))
	// Output: Hola, Juan
}

func ExampleHello_portuguese_brazil() {
	fmt.Println(hello.Hello("pt-br", "João"))
	// Output: Olá, João
}

func ExampleHello_frensh() {
	fmt.Println(hello.Hello("fr", "Jean"))
	// Output: Bonjour, Jean
}
