package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for{
		input, err := reader.ReadString('\n')
		input = strings.Trim(input, "\r\n")
		if err != nil {
			fmt.Println(err)
		}
		conn.Write([]byte(input))
	}

}