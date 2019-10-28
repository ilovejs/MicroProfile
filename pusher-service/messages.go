package main

import (
	"time"
)

const (
	KindProfileCreated = iota + 1
)

type ProfileCreatedMessage struct {
	Kind      uint32    `json:"kind"`
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func newProfileCreatedMessage(id string, body string) *ProfileCreatedMessage {
	return &ProfileCreatedMessage{
		Kind: KindProfileCreated,
		ID:   id,
		Body: body,
	}
}
