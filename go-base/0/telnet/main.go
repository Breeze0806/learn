package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("telnet ip port")
		return
	}
	addr := os.Args[1] + ":" + os.Args[2]
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Println("connct to", addr, "fail. reason:", err)
		return
	}
	conn.Close()
	fmt.Println("connct to", conn.RemoteAddr(), "from", conn.LocalAddr(), "success")
}
