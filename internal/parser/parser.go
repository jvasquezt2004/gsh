package parser

func ParseCommand(tokens []string) (string, []string) {
	return tokens[0], tokens[1:]
}
