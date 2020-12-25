package commands

import (
	s "strings"

	"../command"
)

type deleteCommand struct {
	str, symbol string
}

func (d *deleteCommand) Execute(loop command.Handler) {
	delete := s.Replace(d.str, d.symbol, "", -1)
	loop.Post(&PrintCommand{Arg: delete})
}
