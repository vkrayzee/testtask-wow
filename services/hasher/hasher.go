package hasher

import (
	"strings"

	hc "github.com/catalinc/hashcash"
	"github.com/vkrayzee/testtask-wow/services"
)

type Hasher struct {
	hasher *hc.Hash
}

type HasherConfig struct {
	Bits       int
	SaltLength int
	Extra      string
}

const (
	DefaultBits       = 20
	DefaultSaltLength = 8
	DefaultExtra      = ""
)

// New hasher
func New(config *HasherConfig) services.Hasher {
	// expand default config
	if config == nil {
		config = &HasherConfig{}
	}

	if config.Bits == 0 {
		config.Bits = DefaultBits
	}

	if config.SaltLength == 0 {
		config.SaltLength = DefaultSaltLength
	}

	if config.Extra == "" {
		config.Extra = DefaultExtra
	}

	return &Hasher{
		hasher: hc.New(uint(config.Bits), uint(config.SaltLength), config.Extra),
	}
}

// Check message
func (h *Hasher) Check(message string, challenge string) bool {
	valid := h.hasher.Check(message)
	if !valid {
		return false
	}

	fields := strings.Split(message, ":")

	if len(fields) != 7 {
		return false
	}

	resource := fields[3]

	return resource == challenge
}

// Mint a new stamp
func (h *Hasher) Mint(challenge string) (string, error) {
	return h.hasher.Mint(challenge)
}
