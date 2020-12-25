package commands

import (
	"../command"
)

type CatCommand struct {
	Arg1, Arg2 string
}

func (c *CatCommand) Execute(loop command.Handler) {
	concat := c.Arg1 + c.Arg2
	loop.Post(&PrintCommand{Arg: concat})
}
