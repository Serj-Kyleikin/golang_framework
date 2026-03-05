package subscriptions

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionModel struct {
	ID          uuid.UUID  `db:"-" dbret:"id"`
	ServiceName string     `db:"service_name" dbret:"service_name"`
	Price       int        `db:"price" dbret:"price"`
	UserID      uuid.UUID  `db:"user_id" dbret:"user_id"`
	StartDate   time.Time  `db:"start_date" dbret:"start_date"`
	EndDate     *time.Time `db:"end_date" dbret:"end_date"`

	CreatedAt time.Time `db:"-" dbret:"created_at"`
	UpdatedAt time.Time `db:"-" dbret:"updated_at"`
}
