package commands

import (
	"../command"
)

type ReverseCommand struct {
	Arg string
}

func (r *ReverseCommand) Execute(loop command.Handler) {
	var reverse string
	for _, v := range r.Arg {
		reverse = string(v) + reverse
	}
	loop.Post(&PrintCommand{Arg: reverse})
}
