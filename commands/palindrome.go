package commands

import (
	"../command"
)

type PalindromeCommand struct {
	Arg string
}

func (p *PalindromeCommand) Execute(loop command.Handler) {
	var reverse string
	for _, v := range p.Arg {
		reverse = string(v) + reverse
	}
	palindrome := p.Arg + reverse
	loop.Post(&PrintCommand{Arg: palindrome})
}
