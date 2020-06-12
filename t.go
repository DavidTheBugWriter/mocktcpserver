package main

//https://play.golang.org/p/nC3LpZYAVpR
import (
	"fmt"
	"net"
	"strings"
)

//MockClient mocks a client
type MockClient struct {
	conn net.Conn
}

//NewMockClient create a MockClient
func NewMockClient(conn net.Conn) *MockClient {
	mc := new(MockClient)
	mc.conn = conn
	return mc
}

//MakeFakeConn fakes a net.Conn by using a pipe
func MakeFakeConn() (net.Conn, net.Conn) {
	server, client := net.Pipe()
	return server, client
}

//ReadAll builds a string from the pipe output.
//the read needs to loop in case in missed anything
//in the pipe that stalls the read
func ReadAll(c net.Conn) (string, error) {
	var str strings.Builder
	var readcount int = 0
	var buf = make([]byte, 100)
	for {
		n, err := c.Read(buf)
		str.WriteString(string(buf))
		if err != nil {
			return str.String(), err
		}
		if err == nil {
			return str.String(), err
		}
		readcount += n
	}
}

func main() {
	talk2server, talk2client := MakeFakeConn()

	//For demo simplicity we handle a single client's messages only.
	go handler(talk2client)
	client(talk2server)

	for { //loop forever
	}
}

func handler(c net.Conn) {
	c.Write([]byte("Hi client! Server here...")) //write back to client
	defer c.Close()
	line, _ := ReadAll(c)
	fmt.Println("client said to server:", line)
}

func client(c net.Conn) {
	line, err := ReadAll(c)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("client received from server: ", line)
	c.Write([]byte("hello Server this is my reply."))
}
