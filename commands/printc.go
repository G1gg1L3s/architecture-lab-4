package commands

import (
	s "strings"

	"../command"
)

type printcCommand struct {
	count  int
	symbol string
}

func (p *printcCommand) Execute(loop command.Handler) {
	printc := s.Repeat(p.symbol, p.count)
	loop.Post(&PrintCommand{Arg: printc})
}
