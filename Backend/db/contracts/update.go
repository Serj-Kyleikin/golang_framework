package contracts

import "context"

type Updater[T any, ID comparable] interface {
	TableNamer
	Update(ctx context.Context, id ID, patch any) (T, error)
}
