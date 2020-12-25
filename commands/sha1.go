package commands

import (
	"crypto/sha1"
    "encoding/hex"
	"../command"
)

type sha1Command struct {
	arg string
}

func (s * sha1Command) Execute(loop command.Handler) {
	h := sha1.New()
	h.Write([]byte(s.arg))
	sha1 := hex.EncodeToString(h.Sum(nil))
	loop.Post(&printCommand{arg: sha1})
} 
	