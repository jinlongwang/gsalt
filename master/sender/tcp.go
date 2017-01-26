package sender

import (
	"fmt"
	"net"
	"github.com/gsalt/master/g"
	"bufio"
	"strings"
	"log"
	"github.com/gsalt/master/channel"
)

func InitTcpServer() {
	var tcpAddr *net.TCPAddr
	g.ConnMap = make(map[string]*net.TCPConn)
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1989") //TODO read from config

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		// 新连接加入map
		g.ConnMap["abc"] = tcpConn
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		args := strings.Split(message, "$")
		switch  args[0]{
		case "auth":
			auth(conn, args[1])
		default:
			get(args[1])
		}
		//fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
	}
}

func auth(conn *net.TCPConn, minion_key string) {
	//println(minion_key)
	channel.UserChanel.New(minion_key, conn)
}

func get(key string)  {
	conn := channel.UserChanel.Get(key)
	log.Println("*******")
	log.Println(conn.RemoteAddr().String())
	log.Println("*******")
}