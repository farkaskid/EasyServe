package handler

import (
	"fmt"
	"net"
	"strings"
	"os"
)

type EchoRequestHandler struct{}

func (handler EchoRequestHandler) Handle(connection net.Conn) {
	defer connection.Close()
	
	buffer := make([]byte, 1024)

	connection.Read(buffer)

	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to get the current working directory")

		return
	}

	file, err := os.Open(pwd + "/" + extractPath(buffer))

	if err != nil {
		fmt.Println("Failed to locate the file by path: " + pwd)

		return
	}

	fileStat, err := file.Stat()

	if err != nil {
		fmt.Println("Failed to get file size")

		return
	}

	buffer = make([]byte, fileStat.Size())

	file.Read(buffer)

	connection.Write([]byte("HTTP/1.1 200 OK\n" + string(buffer)))
}

func extractPath(data []byte) string {
	return strings.Split(string(data), " ")[1]
}
