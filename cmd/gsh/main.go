package main

import (
	"os"

	"github.com/jvasquezt2004/gsh/internal/lexer"
)

func main() {
	for {
		tokens, err := lexer.ProcessInput()

		if err != nil {
			continue
		}

		if len(tokens) > 0 && tokens[0] == "exit" {
			os.Exit(0)

		}
	}
}
