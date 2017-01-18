package main

import (
	"github.com/gsalt/minion/receiver"
)

func main()  {
	go receiver.InitTcpClient()
	select {

	}
}
