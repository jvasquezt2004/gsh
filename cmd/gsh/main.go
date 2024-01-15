package main

import (
	"fmt"

	"github.com/jvasquezt2004/gsh/internal/lexer"
	"github.com/jvasquezt2004/gsh/internal/parser"
)

func main() {
	for {
		tokens, err := lexer.ProcessInput()

		if err != nil {
			continue
		}

		command, args := parser.ParseCommand(tokens)

		fmt.Println(command)
		fmt.Println(args)
	}
}
