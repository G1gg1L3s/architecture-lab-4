package commands

import (
	s "strings"

	"../command"
)

type DeleteCommand struct {
	Str, Symbol string
}

func (d *DeleteCommand) Execute(loop command.Handler) {
	delete := s.Replace(d.Str, d.Symbol, "", -1)
	loop.Post(&PrintCommand{Arg: delete})
}
