package main

//https://play.golang.org/p/nC3LpZYAVpR
//https://play.golang.org/p/Zhe3ipWDXsd
import (
	"bytes"
	"fmt"
	"net"
)

var greet = "Hello"

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
	var buf = make([]byte, len(greet))
	for {
		_, err := c.Read(buf)
		newstr := bytes.Trim(buf, "\x00")
		if err != nil {
			return string(newstr), err
		}
		if err == nil {
			return string(newstr), err
		}
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

	c.Write([]byte(greet)) //write back to client
	defer c.Close()
	line, _ := ReadAll(c)
	fmt.Println("client said to server:", line)
	if line == greet {
		fmt.Println("same:", line)
	} else {
		fmt.Println("not same:", line)
	}
}

func client(c net.Conn) {
	line, err := ReadAll(c)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("client received from server: ", line)
	c.Write([]byte(greet))
}
