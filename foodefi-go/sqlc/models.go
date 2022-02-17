// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Blockchains struct {
	ID        int64          `json:"id"`
	Name      sql.NullString `json:"name"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

type EventFields struct {
	ID        int64         `json:"id"`
	EventID   sql.NullInt64 `json:"event_id"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Value     string        `json:"value"`
	Recorder  string        `json:"recorder"`
	CreatedAt sql.NullTime  `json:"created_at"`
}

type Events struct {
	ID           int64         `json:"id"`
	BlockchainID sql.NullInt64 `json:"blockchain_id"`
	BlockNumber  int64         `json:"block_number"`
	EventName    string        `json:"event_name"`
}

type Users struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
