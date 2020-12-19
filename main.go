package main

import "fmt"

import (
	engine "./eventLoop"
	. "./command"
)

func main() {
	ev := new(engine.EventLoop)
	ev.Start()

	ev.Post(CommandFunc(func(h Handler) {
		fmt.Println("1")
	}))
	ev.Post(CommandFunc(func(h Handler) {
		fmt.Println("2")
	}))
	ev.Post(CommandFunc(func(h Handler) {
		fmt.Println("3")
	}))
	ev.AwaitFinish()
}