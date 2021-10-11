package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	ln := scanner.Text()
	// 	fmt.Println(ln)
	// 	fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	// 	if ln == "!DISCONNECT" {
	// 		break
	// 	}
	// }
	
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf))
		conn.Write(buf)
	}

	defer conn.Close()
	fmt.Sprintf("[CLOSED CONNECTION]: %s\n", conn.RemoteAddr())
}

func main() {
	fmt.Println("Start server...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Sprintf("[ERROR]: %s", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
		fmt.Sprintf("[ERROR]: %s", err)
		}
		fmt.Sprintf("[NEW CONNECTION]: %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}	
}