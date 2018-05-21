package main

import (
	"fmt"
	"tcpServer/server"
	"tcpServer/server/handler"
)

func main() {
	port, handler := ":5000", &handler.EchoRequestHandler{}

	fmt.Println("TCP server started at port", port)

	server.Start(port, handler)
}
