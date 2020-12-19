package eventLoop

import (
	queue "../commandQueue"
	. "../command"
)

type EventLoop struct {
	cmdQueue queue.CommandQueue
	exit chan struct{}
	finish bool
}

func (loop *EventLoop) Start()  {
	loop.cmdQueue = queue.InitCommandQueue()
	loop.exit = make(chan struct{})
	loop.finish = false
	go func() {
		for !loop.finish {
			cmd := loop.cmdQueue.Pull()
			cmd.Execute(loop)
		}
		loop.exit <- struct{}{}
	}()
}

func (loop *EventLoop) AwaitFinish() {
	loop.Post(CommandFunc(func( Handler) {
		loop.finish = true
	}))
	<- loop.exit
}

func (loop *EventLoop) Post(cmd Command) {
	loop.cmdQueue.Push(cmd)
}