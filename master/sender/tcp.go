package sender

import (
	"fmt"
	"net"
	"github.com/gsalt/master/g"
	"bufio"
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
		g.ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
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
		fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
		// 这里返回消息改为了广播
		//boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}

