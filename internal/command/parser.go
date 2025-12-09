package command

import (
	"strings"
)

func Parse(command string) Command {
	var tokens []string = tokenizer(command)
	if len(tokens) == 0 {
		return Command{}
	}
	parsed := parser(tokens)

	return parsed
}

func tokenizer(command string) []string {
	var tokens []string
	tokens = strings.Split(command, " ")

	return tokens
}

func parser(tokens []string) Command {
	var command Command

	command.base = tokens[0]
	command.raw = strings.Join(tokens, " ")
	tokens = tokens[1:]

	for _, token := range tokens {
		command.args = append(command.args, token)
	}

	return command
}
