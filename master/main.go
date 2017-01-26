package main

import (
	"github.com/gsalt/master/sender"
	"log"
	//"github.com/gsalt/master/cmd"
	//"time"
	//"github.com/gsalt/master/g"
	//"net"
	"github.com/gsalt/master/channel"
	"github.com/gsalt/master/api"
)

func main()  {
	log.Print("start server")

	channel.UserChanel = channel.NewChannel()
	go sender.InitTcpServer()
	go api.InitApiServer()
	select {

	}

}