package main

import (
	"log"
	"net"
	"os"
)

func handle_client(client net.Conn, target net.Conn) {
	defer client.Close()

	buff := make([]byte, 4096)
	for {
		n, err := client.Read(buff)
		if err != nil {
			log.Panic(err)
			return
		}
		target.Write(buff[0:n])
		log.Printf("Send new packet \033[032m%s\033[0m", string(buff[:n]))
	}
}

func main() {
	LISTEN_PORT := os.Getenv("LISTEN_PORT")

	CONNECT_ADDR := os.Getenv("CONNECT_ADDR")
	CONNECT_PORT := os.Getenv("CONNECT_PORT")

	if LISTEN_PORT == "" {
		log.Panic("LISTEN_PORT is not set.")
		return
	}
	if CONNECT_ADDR == "" {
		log.Panic("CONNECT_ADDR is not set.")
		return
	}
	if CONNECT_PORT == "" {
		log.Panic("CONNECT_PORT is not set.")
		return
	}

	log.Printf("Listen on :%s\n", LISTEN_PORT)
	log.Printf("Connect to %s:%s\n", CONNECT_ADDR, CONNECT_PORT)

	listener, err := net.Listen("tcp", ":"+LISTEN_PORT)
	if err != nil {
		log.Panic(err)
		return
	}
	client, err := net.Dial("tcp", CONNECT_ADDR+":"+CONNECT_PORT)
	if err != nil {
		log.Panic(err)
		return
	}
	for {
		new_client, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handle_client(new_client, client)
		go handle_client(client, new_client)
	}
}
