package main

import (
	"bufio"
	"fmt"
	"os"

	engine "./eventLoop"

	. "./commands"
	. "./parser"
)

func main() {
	ev := new(engine.EventLoop)
	ev.Start()

	if len(os.Args) != 2 {
		ev.Post(&ErrorCommand{Msg: fmt.Sprintf("Usage: %s <file>\n", os.Args[0])})
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			ev.Post(&ErrorCommand{Msg: err.Error()})
		} else {
			reader := bufio.NewReader(file)
			parser := NewParser(reader)
			cmnds := parser.Parse()
			for _, cmd := range cmnds {
				ev.Post(cmd)
			}
		}

	}

	// ev.Post(CommandFunc(func(h Handler) {
	// 	fmt.Println("1")
	// }))
	// ev.Post(CommandFunc(func(h Handler) {
	// 	fmt.Println("2")
	// }))
	// ev.Post(CommandFunc(func(h Handler) {
	// 	fmt.Println("3")
	// }))
	ev.AwaitFinish()
}
