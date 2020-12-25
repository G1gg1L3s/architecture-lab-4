package commands

import (
	"../command"
)

type palindromeCommand struct {
	arg string
}

func (p *palindromeCommand) Execute(loop command.Handler) {
	var reverse string
	for _,v := range p.arg {
		reverse = string(v) + reverse
	}
	palindrome := p.arg + reverse
	loop.Post(&printCommand{arg: palindrome})
}