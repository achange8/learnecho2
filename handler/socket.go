package handler

import (
	"bufio"
	"fmt"
	"net"

	"github.com/labstack/echo"
)

func Socket(c echo.Context) error {
	server, err := net.Listen("tcp", ":8082")
	if err != nil {
		return err
	}

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			return err

		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
	}
	conn.Close()
}
