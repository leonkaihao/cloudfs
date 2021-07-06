package repository

import "context"

type Actions interface {
	Size() int
	Items(ctx context.Context, offset int, size int) ([]ActionItem, error)
}
