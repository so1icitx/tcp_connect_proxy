package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// starts tcp listener on port 8080
	tcpListener, err := net.ListenTCP("tcp", &net.TCPAddr{nil, 8080, ""})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// accepts client connection
		clientTcpconn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}

		// goroutine that handles each client connection concurrently
		go func(clientTcpconn net.Conn) {
			// reads client CONNECT request and stores into data byte slice
			data := make([]byte, 1024)
			clientTcpconn.Read(data)

			// extracts destination domain from CONNECT request line
			var p int
			for data[p] != ':' {
				p++
			}
			domain := string(data[8:p])

			// dns lookup for the extracted domain
			addrs, err := net.LookupHost(domain)
			if err != nil {
				log.Println(err)
			}

			// connects to the clients requested endpoint + signaling the client to prepare to send/receive data
			serverConn, err := net.Dial("tcp", addrs[0]+":443")
			if err != nil {
				log.Println(err)
			}
			clientTcpconn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))

			// two go routines that bidirectionally forward data between client and server
			go func(c net.Conn, s net.Conn) {
				_, err = io.Copy(c, s)
				if err != nil {
					log.Println(err)
				}
			}(clientTcpconn, serverConn)

			go func(c net.Conn, s net.Conn) {
				_, err = io.Copy(s, c)
				if err != nil {
					log.Println(err)
				}
			}(clientTcpconn, serverConn)
		}(clientTcpconn)
	}
}
