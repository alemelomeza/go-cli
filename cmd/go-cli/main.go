package main

import (
	"flag"
	"fmt"

	"github.com/alemelomeza/go-cli/pkg/hello"
)

func main() {
	langPtr := flag.String("lang", "", "Language")
	namePtr := flag.String("name", "", "Name")
	flag.Usage = usage
	flag.Parse()
	fmt.Println(hello.Hello(*langPtr, *namePtr))
}

func usage() {
	msg := fmt.Sprintf("usage: %s [-lang (en | es | pt-br | fr)] [-name <name>]\n\nIt's a simple cli.\n", "go-cli")
	fmt.Println(msg)
	flag.PrintDefaults()
}
