package commands

import (
	s "strconv"

	"../command"
)

type AddCommand struct {
	Arg1, Arg2 int
}

func (a *AddCommand) Execute(loop command.Handler) {
	add := a.Arg1 + a.Arg2
	loop.Post(&PrintCommand{Arg: s.Itoa(add)})
}
