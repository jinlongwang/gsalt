package api

import (
	"os"
	"net/http"
	"github.com/CodisLabs/codis/pkg/utils/log"
)

func InitApiServer()  {
	http.HandleFunc("/channel/get", ChannelHandler)

	if err := http.ListenAndServe("127.0.0.1:7001", nil); err != nil {
		log.Println("start http api server")
		os.Exit(-1)
	}
}
