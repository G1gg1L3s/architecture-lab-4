package commands

import (
	"../command"
)

type ReverseCommand struct {
	arg string
}

func (r *ReverseCommand) Execute(loop command.Handler) {
	var reverse string
	for _, v := range r.arg {
		reverse = string(v) + reverse
	}
	loop.Post(&PrintCommand{Arg: reverse})
}
