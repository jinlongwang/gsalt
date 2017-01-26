package channel

import (
	"net"
	"sync"
	"log"
)

var (
	UserChanel *Channel
)

func NewChannel() *Channel{
	return &Channel{
		Data: make(map[string]*net.TCPConn),
		mutex: &sync.RWMutex{},
	}
}

type Channel struct {
	Data  map[string]*net.TCPConn
	mutex *sync.RWMutex
}

func(c *Channel) New(key string, conn *net.TCPConn)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Data[key] = conn
	log.Println("set key")
	log.Println(c.Data)
}

func(c *Channel) Get(key string) (*net.TCPConn){
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	conn := c.Data[key]
	return conn
}





