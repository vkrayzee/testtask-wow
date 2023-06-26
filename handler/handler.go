package handler

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/vkrayzee/testtask-wow/services"
	"github.com/vkrayzee/testtask-wow/utils"
)

type Handler struct {
	hasher            services.Hasher
	quoteDB           services.QuoteDB
	connectionTimeout time.Duration
}

type Config struct {
	Hasher            services.Hasher
	QuoteDB           services.QuoteDB
	ConnectionTimeout time.Duration
}

// New handler
func New(config *Config) *Handler {
	// expand default config
	if config == nil {
		log.Fatal("QuotesDB is required")
	}

	if config.QuoteDB == nil {
		log.Fatal("QuotesDB is required")
	}

	if config.ConnectionTimeout == 0 {
		config.ConnectionTimeout = 10 * time.Second
	}

	return &Handler{
		hasher:            config.Hasher,
		quoteDB:           config.QuoteDB,
		connectionTimeout: config.ConnectionTimeout,
	}
}

// Handle incoming connection
func (h *Handler) Handle(conn net.Conn) {
	// close connection
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(h.connectionTimeout))

	// print connection info (ip, port)
	fmt.Printf("Connection from %v\n", conn.RemoteAddr())

	challenge := ""

	if h.hasher != nil {
		// generate challenge
		challenge = utils.GenerateChallenge()
	}

	// write challenge to response
	conn.Write([]byte(challenge))

	// incoming request
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	message := string(buffer[:n])

	valid := true

	if h.hasher != nil {
		// check a stamp
		valid = h.hasher.Check(message, challenge)
	}

	if !valid {
		log.Println("Invalid")
		return
	}

	// write data to response
	responseStr := h.quoteDB.GetRandomQuote()
	conn.Write([]byte(responseStr))
}
