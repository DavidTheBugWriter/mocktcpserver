package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

//MakeFakeConn fakes a net.Conn by using a pipe
func MakeFakeConn() (net.Conn, net.Conn) {
	server, client := net.Pipe()
	return server, client
}

func main() {
	//create fake sockets from net.Pipe
	talk2server, talk2client := MakeFakeConn()

	//For demo simplicity we handle a single client's messages only.
	go client(talk2server)
	go handler(talk2client)
	for { //loop forever
	}
}

func handler(c net.Conn) {
	message, _ := bufio.NewReader(c).ReadString('\n')
	fmt.Println("message from client:", message)
	c.Write([]byte("Hi client! Server here...")) //write back to client
	c.Close()
}

func client(c net.Conn) {
	c.Write([]byte("hello Server.\r\n"))
	lr := io.LimitReader(c, 512) //DDoS protection
	buffread := bufio.NewReader(lr)
	lines, _, _ := buffread.ReadLine()
	fmt.Println("received from server: ", string(lines))
}
