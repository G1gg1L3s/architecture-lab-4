package commands

import (
	"../command"
)

type palindromeCommand struct {
	Arg string
}

func (p *palindromeCommand) Execute(loop command.Handler) {
	var reverse string
	for _, v := range p.Arg {
		reverse = string(v) + reverse
	}
	palindrome := p.Arg + reverse
	loop.Post(&PrintCommand{Arg: palindrome})
}
