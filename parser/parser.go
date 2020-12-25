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

type tokenType int

const (
	num tokenType = iota
	str
	sym
)

type parseRule struct {
	constructor func(args []interface{}) Command
	rule        []tokenType
}

var parseTable = map[string]parseRule{
	"print":     {rule: []tokenType{str}, constructor: printConst},
	"printc":    {rule: []tokenType{num, sym}, constructor: printcConst},
	"add":       {rule: []tokenType{num, num}, constructor: addConst},
	"palindrom": {rule: []tokenType{str}, constructor: palindromeConst},
	"split":     {rule: []tokenType{str, sym}, constructor: splitConst},
	"reverse":   {rule: []tokenType{str}, constructor: reverseConst},
	"sha1":      {rule: []tokenType{str}, constructor: sha1Const},
	"delete":    {rule: []tokenType{str, sym}, constructor: deleteConst},
	"cat":       {rule: []tokenType{str, str}, constructor: catConst},
}

func printConst(args []interface{}) Command {
	arg := args[0].(string)
	return &PrintCommand{Arg: arg}
}

func printcConst(args []interface{}) Command {
	count := args[0].(int)
	sym := args[1].(string)
	return &PrintcCommand{Count: count, Symbol: sym}
}

func addConst(args []interface{}) Command {
	arg1 := args[0].(int)
	arg2 := args[1].(int)
	return &AddCommand{Arg1: arg1, Arg2: arg2}
}

func palindromeConst(args []interface{}) Command {
	arg := args[0].(string)
	return &PalindromeCommand{Arg: arg}
}

func splitConst(args []interface{}) Command {
	str := args[0].(string)
	sep := args[1].(string)
	return &SplitCommand{Str: str, Sep: sep}
}

func reverseConst(args []interface{}) Command {
	arg := args[0].(string)
	return &ReverseCommand{Arg: arg}
}

func sha1Const(args []interface{}) Command {
	arg := args[0].(string)
	return &Sha1Command{Arg: arg}
}

func deleteConst(args []interface{}) Command {
	str := args[0].(string)
	sym := args[1].(string)
	return &DeleteCommand{Str: str, Symbol: sym}
}

func catConst(args []interface{}) Command {
	arg1 := args[0].(string)
	arg2 := args[1].(string)
	return &CatCommand{Arg1: arg1, Arg2: arg2}
}

type Parser struct {
	input *bufio.Scanner
	line  int
}

func (p *Parser) parseLine(line string) Command {
	parts := strings.Fields(line)
	cmd := parts[0]

	rule, ok := parseTable[cmd]
	if !ok {
		return p.errorCommand(fmt.Sprintf("Unknown command: %s", cmd))
	}
	if len(rule.rule) == 1 && rule.rule[0] == str {
		// The command receives only one argument - the whole string
		start := len(cmd)
		rest := strings.Trim(line[start:], " ")
		return rule.constructor([]interface{}{rest})
	}
	parts = parts[1:]
	if len(parts) != len(rule.rule) {
		return p.errorCommand(fmt.Sprintf("Error in %s: expect %d args, got %d",
			cmd,
			len(rule.rule),
			len(parts)),
		)
	}
	return p.mathArgs(cmd, rule, parts)
}

func (p *Parser) mathArgs(cmd string, rule parseRule, args []string) Command {
	res := []interface{}{}
	for i := range rule.rule {
		switch rule.rule[i] {
		case num:
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return p.errorCommand(fmt.Sprintf("Error in %s while parsing number: %s", cmd, err.Error()))
			}
			res = append(res, n)
		case str:
			res = append(res, args[i])
		case sym:
			sym := args[i]
			if len([]rune(sym)) != 1 {
				return p.errorCommand(fmt.Sprintf("Error in %s: got string instead of single character", cmd))
			}
			res = append(res, args[i])
		}
	}
	return rule.constructor(res)
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
