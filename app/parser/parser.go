package parser

import (
	"errors"
	"fmt"
	"strings"
)

type Parser struct {
	Name        string
	Args        []string
	Redirection Redirection
	Filepath    string
	tokens      []string
}

type Redirection int

const (
	NoRedirection Redirection = iota
	Stdout
	StdoutAppend
	Stderr
	StderrAppend
)

// NewParser creates a new parser and immediately parses the input
func NewParser(input string) *Parser {
	p := &Parser{}
	p.createTokens(input)
	p.Name = p.tokens[0] // Simple example, adjust as needed
	p.Redirection = NoRedirection
	p.Args = []string{}
	if len(p.tokens) < 2 {
		return p
	}
	var index int
	var err error
	p.Redirection, index, err = p.getRedirectionToken()
	if err != nil {
		fmt.Println(err)
	}

	if p.Redirection != NoRedirection {
		p.Filepath = p.tokens[index+1]
		p.Args = p.tokens[1:index]
	} else {
		p.Args = p.tokens[1:]
	}

	return p
}

func (p *Parser) createTokens(input string) {
	var currentToken strings.Builder
	inSingleQuote := false
	inDoubleQuote := false
	escapeNext := false

	for i, char := range input {
		switch {
		case escapeNext:
			currentToken.WriteRune(char)
			escapeNext = false
		case char == '\\' && !inDoubleQuote && !inSingleQuote:
			escapeNext = true
		case char == '\\' && inDoubleQuote:
			if input[i+1] == '"' || input[i+1] == '\\' || input[i+1] == '$' {
				escapeNext = true
			} else {
				currentToken.WriteRune(char)
			}
		case char == '\'' && !inDoubleQuote:
			inSingleQuote = !inSingleQuote
		case char == '"' && !inSingleQuote:
			inDoubleQuote = !inDoubleQuote
		case char == ' ' && !inSingleQuote && !inDoubleQuote:
			if currentToken.Len() > 0 {
				p.tokens = append(p.tokens, currentToken.String())
				currentToken.Reset()
			}
		case char == '\n' && !inSingleQuote && !inDoubleQuote:
			continue
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		p.tokens = append(p.tokens, currentToken.String())
	}
}

func (p *Parser) getRedirectionToken() (Redirection, int, error) {
	// Implement redirection parsing
	redirection := NoRedirection
	redirectionIndex := len(p.tokens)

	for i, token := range p.tokens {
		if token == ">" || token == "1>" {
			redirection = Stdout
			redirectionIndex = i
			break
		} else if token == ">>" {
			redirection = StdoutAppend
			redirectionIndex = i
			break
		} else if token == "2>" {
			redirection = Stderr
			redirectionIndex = i
			break
		}
	}

	if redirectionIndex == len(p.tokens)-1 && redirection != NoRedirection {
		return NoRedirection, redirectionIndex, errors.New("no file specified for redirection")
	}

	return redirection, redirectionIndex, nil
}
