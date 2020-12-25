package commands

import (
	s "strings"

	"../command"
)

type PrintcCommand struct {
	Count  int
	Symbol string
}

func (p *PrintcCommand) Execute(loop command.Handler) {
	printc := s.Repeat(p.Symbol, p.Count)
	loop.Post(&PrintCommand{Arg: printc})
}
