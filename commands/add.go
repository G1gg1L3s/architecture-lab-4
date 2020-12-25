package commands

import (
	s "strconv"
	"../command"
)

type addCommand struct {
	arg1, arg2 int
}

func (a *addCommand) Execute(loop command.Handler) {
	add := a.arg1 + a.arg2
	loop.Post(&printCommand{arg: s.Itoa(add)})
} 
	