package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(chan string)
	go clientWriter(conn, ch)
	ip := conn.RemoteAddr().String()
	log.Println("You are ip : ", ip)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		//fmt.Fprintf(conn, msg)
		log.Println(msg)
	}

}
func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
	fmt.Println("hello")
}
