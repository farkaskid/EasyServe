package server

import (
	"fmt"
	"net"
)

type RequestHandler interface {
	Handle(connection net.Conn)
}

func Start(port string, handler RequestHandler) {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic("Error occured")
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error occured", err)
		}

		go handler.Handle(conn)
	}
}
