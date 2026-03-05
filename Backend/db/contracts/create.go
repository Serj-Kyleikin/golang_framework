package contracts

import "context"

type Creator[T any] interface {
	Create(ctx context.Context, entity T) (T, error)
}
