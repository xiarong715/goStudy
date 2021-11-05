package main

import (
	"bufio"
	"fmt"
	"net"
)

func proccess(conn net.Conn) {
	var buf [512]byte
	reader := bufio.NewReader(conn)
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Println("read failed")
		return
	}
	conn.Write(buf[:n])
}

func main() {
	fmt.Println("server start ...")
	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println("listen failed: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed: ", err)
			continue
		}
		go proccess(conn)
	}
}
