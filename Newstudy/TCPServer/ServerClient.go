package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error, info string) (res bool) {
	if err != nil {
		fmt.Println(info + " " + err.Error())
		return false
	}
	return true
}

////////////////////////////////////
//服务器端接收数据线程
//参数：
//		数据连接 conn
//		通讯通道 messages
////////////////////////////////////
func Handler(conn net.Conn, messages chan string) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			break
		}
		if length > 0 {
			buf[length] = 0
		}
		reciveStr := string(buf[0:length])
		messages <- reciveStr
	}
}

///////////////////////////////
//服务器发送数据的线程
//参数
//		连接字典 conns
//		数据通道 messages
///////////////////////////////
func echoHandler(conns *map[string]net.Conn, messages chan string) {

	for {
		msg := <-messages
		fmt.Println(msg)
		for key, value := range *conns {
			fmt.Println("Connection is connected from ...", key)
			_, err := value.Write([]byte(msg))
			if err != nil {
				fmt.Println(err.Error())
				delete(*conns, key)
			}
		}
	}
}

func StartServer(port string) {
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)
	go echoHandler(&conns, messages)
	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		fmt.Println("Accepting ...")
		conns[conn.RemoteAddr().String()] = conn
		//启动新线程
		go Handler(conn, messages)
	}
}

///////////////////
//客户端发送线程
//参数
//		发送连接 conn
//
//
///////////////////
func chatSend(conn net.Conn) {
	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("ByeBye...")
			conn.Close()
			os.Exit(0)
		}
		lens, err := conn.Write([]byte(username + " Say:::" + input))
		fmt.Println(lens)
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}
	}
}

//////////////////
//客户端启动函数
//参数
//		远程ip地址和端口
//
/////////////////
func StartClient(tcpaddr string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpaddr)
	if checkError(err, "ResolveTCPAddr") == false {
		os.Exit(0)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	go chatSend(conn)
	buff := make([]byte, 1024)
	for {
		length, err := conn.Read(buff)
		if checkError(err, "Connection") == false {
			conn.Close()
			fmt.Println("Server is dead ...ByeBye")
			os.Exit(0)
		}
		fmt.Println(string(buff[0:length]))
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong pare")
		os.Exit(0)
	}
	if os.Args[1] == "server" && len(os.Args) == 3 {
		StartServer(os.Args[2])
	}
	if os.Args[1] == "client" && len(os.Args) == 3 {
		StartClient(os.Args[2])
	}
}
