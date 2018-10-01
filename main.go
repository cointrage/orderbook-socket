package main

// Broadcasting socket server for orderbook messages.

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string // an outgoing message channel

const (
	socketHost = "localhost"
	socketPort = 3009
)

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func main() {
	
	listener, err := net.Listen("tcp", socketHost + ":" + socketPort)
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}

}

func broadcaster() {
	
	clients := make(map[client]bool) // all connected clients
	
	for {
		select {
			case msg := <-messages:
				// broadcast incoming message to all
				for cli := range clients {
					cli <- msg
				}

			case cli := <-entering:
				clients[cli] = true

			case cli := <-leaving:
				delete(clients, cli)
				close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	
	ch := make(chan string) // outgoing client messages

	go clientWriter(conn, ch)

	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- input.Text()
	}

	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg:= range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
