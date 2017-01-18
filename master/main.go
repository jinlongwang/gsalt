package main

import (
	"github.com/gsalt/master/sender"
	"log"
	"github.com/gsalt/master/cmd"
)

func main()  {
	log.Print("start server")
	go sender.InitTcpServer()
	go cmd.InitCmd()
	select {

	}

}