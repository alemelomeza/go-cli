package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseBrazilHelloPrefix = "Olá, "
const frenshHelloPrefix = "Bonjour, "

// Hello say hello in different languages
func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "es":
		prefix = spanishHelloPrefix
	case "pt-br":
		prefix = portugueseBrazilHelloPrefix
	case "fr":
		prefix = frenshHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
