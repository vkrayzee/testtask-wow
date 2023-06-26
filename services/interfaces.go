package services

// QuoteDB is a database of quotes
type QuoteDB interface {
	GetRandomQuote() string
	Close()
}

// Hasher is a service for hashing messages
type Hasher interface {
	Check(message string, challenge string) bool
	Mint(challenge string) (string, error)
}
