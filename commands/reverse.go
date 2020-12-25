package commands

import (
	"../command"
)

type reverseCommand struct {
	arg string
}

func (r *reverseCommand) Execute(loop command.Handler) {
	var reverse string
	for _, v := range r.arg {
		reverse = string(v) + reverse
	}
	loop.Post(&PrintCommand{Arg: reverse})
}
