package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Command interface {
	Execute()
}

type PlugCmd struct{}

func (cmd *PlugCmd) Execute() {
	fmt.Println("Executing...")
}

type Parser struct {
	input *bufio.Scanner
	line  int
}

func (p *Parser) parsePrint(rest string) Command {
	return &PlugCmd{}
}

func (p *Parser) parseSha1(rest string) Command {
	if len(rest) == 0 {
		return p.errorCommand(fmt.Sprintf("Expected string, found nothing"))
	}
	return &PlugCmd{}
}

func (p *Parser) parseLine(line string) Command {
	cmd := strings.Fields(line)[0]
	start := len(cmd)
	rest := strings.Trim(line[start:], " ")
	if cmd == "print" {
		return p.parsePrint(rest)
	} else if cmd == "sha1" {
		return p.parseSha1(rest)
	} else {
		return p.errorCommand(fmt.Sprintf("Unknown command: %s", cmd))
	}
}

func (p *Parser) parse() []Command {
	res := []Command{}
	for p.input.Scan() {
		p.line++
		line := p.input.Text()
		cmd := p.parseLine(line)
		res = append(res, cmd)
	}
	return res
}

func (p *Parser) parseNext() Command {
	p.line++
	p.input.Scan()
	line := p.input.Text()
	return p.parseLine(line)
}

func NewParser(reader io.Reader) Parser {
	return Parser{input: bufio.NewScanner(reader), line: 0}
}

func (p *Parser) errorCommand(msg string) Command {
	msg += fmt.Sprintf(" on line %d", p.line)
	return &PlugCmd{}
}
