package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vkrayzee/testtask-wow/handler"
	"github.com/vkrayzee/testtask-wow/services/hasher"
	"github.com/vkrayzee/testtask-wow/services/quotes"
)

const (
	HOST = "localhost"
	PORT = 8080
	TYPE = "tcp"
)

func main() {
	// get host and port from args
	host := flag.String("host", HOST, "host")
	port := flag.Int("port", PORT, "port")

	flag.Parse()

	// graceful shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	quoteDB := quotes.New()

	// listen
	listen, err := net.ListenTCP(TYPE, &net.TCPAddr{
		IP:   net.ParseIP(*host),
		Port: *port,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// close listen socket
	defer listen.Close()

	fmt.Printf("Server is listening on %v\n", listen.Addr())

	// hasher
	h := hasher.New(nil)

	// handler
	handler := handler.New(&handler.Config{
		Hasher:  h,
		QuoteDB: quoteDB,
	})

	wg := sync.WaitGroup{}

loop:
	for {
		select {
		case <-exit:
			fmt.Println("Server is shutting down...")
			break loop
		default:
			listen.SetDeadline(time.Now().Add(1 * time.Second))

			// accept connection
			conn, err := listen.Accept()
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}
			if err != nil {
				log.Println(err)
				continue
			}

			// handle request
			wg.Add(1)
			go func() {
				defer wg.Done()
				handler.Handle(conn)
			}()
		}
	}

	listen.Close()
	quoteDB.Close()
	wg.Wait()

	fmt.Println("Server is shutted down")
}
