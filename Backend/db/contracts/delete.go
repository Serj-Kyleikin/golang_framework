package contracts

import "context"

type Deleter[ID comparable] interface {
	TableNamer
	Delete(ctx context.Context, id ID) error
}
