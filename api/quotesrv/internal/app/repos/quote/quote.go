package quote

import (
	"context"
	"github.com/google/uuid"
)

type Quote struct {
	ID       uuid.UUID
	BookName string
	Text     string
	Source   int
}

// Store нужен только тут
type Store interface {
	Create(ctx context.Context, q Quote) (*uuid.UUID, error)
	Read(ctx context.Context, uid uuid.UUID) (*Quote, error)
	Delete(ctx context.Context, uid uuid.UUID) error
	SearchQuotes(ctx context.Context, s string) (chan Quote, error)
}

type Quotes struct {
	store Store
}

func NewQuotes(store Store) *Quotes {
	return &Quotes{
		store: store,
	}
}
