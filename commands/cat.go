package commands

import (
	"../command"
)

type catCommand struct {
	arg1, arg2 string
}

func (c *catCommand) Execute(loop command.Handler) {
	concat := c.arg1 + c.arg2
	loop.Post(&PrintCommand{Arg: concat})
}
