package lexer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessInput() ([]string, error) {
	var tokens []string
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		command := scanner.Text()
		words := strings.Fields(command)
		tokens = append(tokens, words...)
		break
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}
