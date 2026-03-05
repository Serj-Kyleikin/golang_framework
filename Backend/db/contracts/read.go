package contracts

import "context"

type Reader[T any, ID comparable] interface {
	TableNamer
	GetByID(ctx context.Context, id ID) (T, error)
	List(ctx context.Context, limit, offset int) ([]T, error)
}
