package commands

import (
	s "strings"

	"../command"
)

type SplitCommand struct {
	Str, Sep string
}

func (sp *SplitCommand) Execute(loop command.Handler) {
	split := s.Split(sp.Str, sp.Sep)
	for _, v := range split {
		loop.Post(&PrintCommand{Arg: v})
	}
}
