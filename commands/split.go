package commands

import (
    s "strings"
	"../command"
)

type splitCommand struct {
	str, sep string
}

func (sp *splitCommand) Execute(loop command.Handler) {
	split := s.Split(sp.str, sp.sep)
	for _,v := range split {
		loop.Post(&printCommand{arg: v})
	}
}