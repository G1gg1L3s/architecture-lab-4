package commands

import (
	"fmt"
	"../command"
)

type printCommand struct {
	arg string
}
	
func (p *printCommand) Execute(loop command.Handler) {
	fmt.Println(p.arg)
} 