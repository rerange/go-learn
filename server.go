package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		go func(c net.Conn) {
			buffer := make([]byte, 512)
			for {
				num, err := c.Read(buffer)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(buffer[:num]))
			}
		}(conn)
	}
}

func handle(c net.Conn) {
	
}