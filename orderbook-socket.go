package main

// Broadcasting socket server for orderbook messages.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"./orderbook"
)

const (
	socketPort = 3009
)

type client chan<- string // an outgoing message channel

type clientMessage struct {
	Client  client
	Message string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan *clientMessage) // all incoming messages
)

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", socketPort))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on port: %d", socketPort)

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

	clients := make(map[client]bool)     // all connected clients
	subscribers := make(map[client]bool) // subscribers for broadcasting

	for {
		select {
		case msg := <-messages:

			// decoding message
			var message orderbook.Message
			if err := json.Unmarshal([]byte((*msg).Message), &message); err != nil {
				log.Printf("could not parse incoming message: %v, %v", err, (*msg).Message)
				continue
			}

			switch message.Type {
			case orderbook.MessageSubscribe:
				// registering client as subscriber
				subscribers[(*msg).Client] = true

			case orderbook.MessageBook:
				// broadcast orderbook message to all subscribers
				for cli := range subscribers {
					cli <- (*msg).Message
				}

			case orderbook.MessageDiff:
				// broadcast orderbook message to all subscribers
				for cli := range subscribers {
					cli <- (*msg).Message
				}

			default:
				log.Printf("unknown message type received: %s", (*msg).Message)
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			delete(subscribers, cli)

			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {

	ch := make(chan string) // outgoing client messages
	clientAddr := conn.RemoteAddr().String()

	go clientWriter(conn, ch)

	// registering client
	log.Printf("new socket client %s connected", clientAddr)
	entering <- ch

	// reading input
	r := bufio.NewReader(conn)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			log.Printf("socket read error: %v", err)
			break
		}

		messages <- &clientMessage{ch, str[:len(str)-1]}
	}

	// deregistering client
	leaving <- ch
	log.Printf("socket client %s disconnected", clientAddr)

	// closing socket connection
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
