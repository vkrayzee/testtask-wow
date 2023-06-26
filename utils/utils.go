package utils

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// GenerateChallenge generates challenge string (sha1 from "some challenge" and timestamp)
func GenerateChallenge() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte("some challenge"+time.Now().String())))
}
