package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8002")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		/*
			n, err := io.Copy(os.Stderr, conn)
			log.Printf("Copied %d bytes; finished with err = %v ", n, err)

		*/
		//go copyToStderr(conn)

		//go copyToStderrSettimeout(conn)

		go proxy(conn)
	}

}

func proxy(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", "google.com:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, conn)
	io.Copy(conn, remote)

}

func copyToStderrSettimeout(conn net.Conn) {
	//n, err := io.Copy(os.Stderr, conn)
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Finished with err = %v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
	//log.Printf("Copied %d bytes; finished with err = %v ", n, err)
}

func copyToStderr(conn net.Conn) {
	n, err := io.Copy(os.Stderr, conn)
	log.Printf("Copied %d bytes; finished with err = %v ", n, err)
}
