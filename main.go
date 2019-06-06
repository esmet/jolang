package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	p := NewParser(string(input))

	ast, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	result, err := ast.eval()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %d\n", result)
}
