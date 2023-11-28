package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	now := time.Now()
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  now.Unix(),
		ExpiredAt: now.Add(duration).Unix(),
	}

	return payload, nil
}

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

func (payload *Payload) Valid() error {
	if time.Unix(payload.ExpiredAt, 0).Before(time.Now()) {
		return ErrExpiredToken
	}

	return nil
}
