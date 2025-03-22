package parser

import (
	"errors"
	"fmt"
	"strings"
)

type Parser struct {
	tokens      []string
	command     string
	args        []string
	redirection Redirection
	filepath    string
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
	p.command = p.tokens[0] // Simple example, adjust as needed
	p.redirection = NoRedirection
	p.args = []string{}
	if len(p.tokens) < 2 {
		return p
	}
	var index int
	var err error
	p.redirection, index, err = p.getRedirectionToken()
	if err != nil {
		fmt.Println(err)
	}

	if p.redirection != NoRedirection {
		p.filepath = p.tokens[index+1]
		p.args = p.tokens[1:index]
	} else {
		p.args = p.tokens[1:]
	}

	return p
}

// GetCommand returns the parsed command
func (p *Parser) GetCommand() string {
	return p.command
}

// GetRedirection returns any redirection found
func (p *Parser) GetRedirection() Redirection {
	return p.redirection
}

func (p *Parser) GetFilepath() string {
	return p.filepath
}

// GetArgs returns command arguments
func (p *Parser) GetArgs() []string {
	return p.args
}

func (p *Parser) createTokens(input string) {
	var tokens []string
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
				p.tokens = append(tokens, currentToken.String())
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
	var i int
	var token string
	var redirection Redirection

	for i, token = range p.tokens {
		if token == ">" || token == "1>" {
			redirection = Stdout
		} else if token == ">>" {
			redirection = StdoutAppend
		} else if token == "2>" {
			redirection = Stderr
		} else {
			redirection = NoRedirection
		}
	}

	if i == len(p.tokens)-1 && redirection != NoRedirection {
		return NoRedirection, i, errors.New("no file specified for redirection")
	}

	return redirection, i, nil
}
