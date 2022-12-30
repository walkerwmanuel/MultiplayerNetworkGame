package main

import (
	"fmt"
	"os"
	"strconv"
	"webgame/server"
)

func main() {
	args := os.Args
	s := server.Server{}
	if len(args) > 1 {
		newPort, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err != nil {
			fmt.Println("bad port provided", os.Args[1])
		} else {
			s.Port = int(newPort)
		}
	}
	s.Start()
}
