package repository

import "github.com/google/uuid"

type ActionItem struct {
	ID uuid.UUID `json:"id"`
}

type HasheItem struct {
	ID uuid.UUID `json:"id"`
}

type WatchItem struct {
	ID uuid.UUID `json:"id"`
}
