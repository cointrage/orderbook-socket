package main

// Broadcasting socket server for orderbook messages.

import (
	"fmt"
	"log"
	"net"
	"github.com/cointrage/orderbook-socket/orderbook"
	"github.com/gogo/protobuf/io"
	"time"
)

const (
	socketPort = 3009
	readLimit  = 100000
	statsLimit = 100000
)

type client chan<- *orderbook.Message // an outgoing message channel

type clientMessage struct {
	Client  client
	Message *orderbook.Message
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

	start := time.Now()
	var counter int64
	var totalsize int64

	for {
		select {
		case msg := <-messages:

			message := msg.Message

			// recording statistics
			counter++
			totalsize += int64(len(message.Data))
			if counter % statsLimit == 0 {
				avgrate := float32(counter)/float32(time.Since(start).Seconds())
				avgsize := float32(totalsize)/float32(counter)
				log.Printf("performing at avg. rate %0.1f msgs/s with avg. message size %0.1f bytes", avgrate, avgsize)

				// reset
				start = time.Now()
				counter = 0
				totalsize = 0
			}

			// handling message
			switch message.Type {
			case orderbook.MessageType_SubscribeMessage:
				// registering client as subscriber
				subscribers[(*msg).Client] = true

			case orderbook.MessageType_OrderBookMessage:
				// broadcast orderbook message to all subscribers
				for cli := range subscribers {
					cli <- message
				}

			case orderbook.MessageType_OrderBookDiffMessage:
				// broadcast orderbook message to all subscribers
				for cli := range subscribers {
					cli <- message
				}

			default:
				log.Printf("unknown message type received: %s", msg.Message)
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

	ch := make(chan *orderbook.Message) // outgoing client messages
	clientAddr := conn.RemoteAddr().String()

	go clientWriter(conn, ch)

	// registering client
	log.Printf("new socket client %s connected", clientAddr)
	entering <- ch

	// reading input
	r := io.NewDelimitedReader(conn, readLimit)
	for {
		msg := &orderbook.Message{}
		if err := r.ReadMsg(msg); err != nil {
			log.Printf("socket read error: %v", err)
			break
		}

		messages <- &clientMessage{ch, msg}
	}

	// deregistering client
	leaving <- ch
	log.Printf("socket client %s disconnected", clientAddr)

	// closing socket connection
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan *orderbook.Message) {

	w := io.NewDelimitedWriter(conn)

	for msg := range ch {
		w.WriteMsg(msg)
	}
}
