package token

import "time"

// Maker is interface for managing tokens
type Maker interface {
	// CreateToken Creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken Verifies if a token is valid or not
	VerifyToken(token string) (*Payload, error)
}
