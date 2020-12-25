package parser

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	. "../command"
	. "../commands"
)

type Parser struct {
	input *bufio.Scanner
	line  int
}

func (p *Parser) parsePrint(rest string) Command {
	return &PrintCommand{Arg: rest}
}

func (p *Parser) parseReverse(rest string) Command {
	return &ReverseCommand{Arg: rest}
}

func (p *Parser) parsePalindrome(rest string) Command {
	return &PalindromeCommand{Arg: rest}
}

func (p *Parser) parsePrintC(rest string) Command {
	parts := strings.Fields(rest)
	if len(parts) != 2 {
		return p.errorCommand(fmt.Sprintf("Error in printc: expect 2 args"))
	}
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return p.errorCommand(fmt.Sprintf("Error parsing number: %s", err.Error()))
	}
	symbol := parts[1]
	return &PrintcCommand{Count: count, Symbol: symbol}
}

func (p *Parser) parseAdd(rest string) Command {
	parts := strings.Fields(rest)
	if len(parts) != 2 {
		return p.errorCommand(fmt.Sprintf("Error in add: expect 2 args"))
	}
	arg1, err := strconv.Atoi(parts[0])
	if err != nil {
		return p.errorCommand(fmt.Sprintf("Error parsing first number: %s", err.Error()))
	}
	arg2, err := strconv.Atoi(parts[1])
	if err != nil {
		return p.errorCommand(fmt.Sprintf("Error parsing second number: %s", err.Error()))
	}
	return &AddCommand{Arg1: arg1, Arg2: arg2}
}

func (p *Parser) parseSha1(rest string) Command {
	if len(rest) == 0 {
		return p.errorCommand(fmt.Sprintf("Expected string, found nothing"))
	}
	return &Sha1Command{Arg: rest}
}

func (p *Parser) parseSplit(rest string) Command {
	parts := strings.Fields(rest)
	if len(parts) != 2 {
		return p.errorCommand(fmt.Sprintf("Error in split: expect 2 args"))
	}
	return &SplitCommand{Str: parts[0], Sep: parts[1]}
}

func (p *Parser) parseDelete(rest string) Command {
	parts := strings.Fields(rest)
	if len(parts) != 2 {
		return p.errorCommand(fmt.Sprintf("Error in delete: expect 2 args"))
	}
	return &DeleteCommand{Str: parts[0], Symbol: parts[1]}
}

func (p *Parser) parseCat(rest string) Command {
	parts := strings.Fields(rest)
	if len(parts) != 2 {
		return p.errorCommand(fmt.Sprintf("Error in cat: expect 2 args"))
	}
	return &CatCommand{Arg1: parts[0], Arg2: parts[1]}
}

func (p *Parser) parseLine(line string) Command {
	cmd := strings.Fields(line)[0]
	start := len(cmd)
	rest := strings.Trim(line[start:], " ")
	switch cmd {
	case "print":
		return p.parsePrint(rest)
	case "add":
		return p.parseAdd(rest)
	case "palindrom":
		return p.parsePalindrome(rest)
	case "split":
		return p.parseSplit(rest)
	case "reverse":
		return p.parseReverse(rest)
	case "sha1":
		return p.parseSha1(rest)
	case "printc":
		return p.parsePrintC(rest)
	case "delete":
		return p.parseDelete(rest)
	case "cat":
		return p.parseCat(rest)
	default:
		return p.errorCommand(fmt.Sprintf("Unknown command: %s", cmd))
	}
}

func (p *Parser) Parse() []Command {
	res := []Command{}
	for p.input.Scan() {
		p.line++
		line := p.input.Text()
		// Skip empty lines
		if len(strings.Trim(line, " \t\n")) == 0 {
			continue
		}
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
	msg += fmt.Sprintf(" on line %d\n", p.line)
	return &ErrorCommand{Msg: msg}
}
