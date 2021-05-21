package hello_test

import (
	"fmt"
	"testing"

	"github.com/alemelomeza/go-cli/pkg/hello"
)

var tt = []struct {
	lang string
	name string
	want string
}{
	{"", "", "Hello, World"},
	{"en", "", "Hello, World"},
	{"en", "John", "Hello, John"},
	{"es", "", "Hola, World"},
	{"es", "Juan", "Hola, Juan"},
	{"pt-br", "", "Olá, World"},
	{"pt-br", "João", "Olá, João"},
	{"fr", "", "Bonjour, World"},
	{"fr", "Jean", "Bonjour, Jean"},
}

func TestHello(t *testing.T) {
	for _, tc := range tt {
		subtest := fmt.Sprintf("%v %v", tc.lang, tc.name)
		t.Run(subtest, func(t *testing.T) {
			if got := hello.Hello(tc.lang, tc.name); got != tc.want {
				t.Errorf("got %q want %q", got, tc.want)
			}
		})
	}
}
