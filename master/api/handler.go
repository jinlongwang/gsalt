package api

import (
	"net/http"
	"encoding/json"
	"io"
	"github.com/gsalt/master/channel"
	"log"
)

func ChannelHandler(rw http.ResponseWriter, r *http.Request)  {
	var(
		result = make(map[string]interface{})
		status = 0

	)
	key := r.URL.Query().Get("key")
	if key == ""{
		res := make(map[string]string)
		for key, value := range channel.UserChanel.Data {
			log.Println("key is:" + key + " value is:" + value.RemoteAddr().String())
			res[key] = value.RemoteAddr().String()
		}
		result["msg"] = res
		result["status"] = status
	}else {
		log.Println(key)
		conn := channel.UserChanel.Get(key)
		log.Println("-----------")
		log.Println(channel.UserChanel.Data["123"].RemoteAddr().String())
		log.Println("-----------")
		log.Println(conn)
		//address := conn.RemoteAddr().String()
		//log.Println(address)
		result["msg"] = 123
		result["status"] = status
	}

	date, _ := json.Marshal(result)
	io.WriteString(rw, string(date))
}
