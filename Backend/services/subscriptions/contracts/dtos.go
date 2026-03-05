package contracts

import (
	"time"

	"github.com/google/uuid"
)

// CREATE
type CreateInputDTO struct {
	ServiceName string
	Price       int
	UserID      uuid.UUID

	StartDate time.Time
	EndDate   *time.Time
}

type CreateOutputDTO struct {
	ID          string `json:"id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`

	StartDate string  `json:"start_date"`
	EndDate   *string `json:"end_date,omitempty"`
}

// READ
type GetInputDTO struct {
	ID uuid.UUID
}

type GetOutputDTO struct {
	ID          string  `json:"id"`
	ServiceName string  `json:"service_name"`
	Price       int     `json:"price"`
	UserID      string  `json:"user_id"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date,omitempty"`
}

// LIST
type ListInputDTO struct {
	Limit  int
	Offset int
}

type ListOutputDTO struct {
	Items  []GetOutputDTO `json:"items"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
}

// UPDATE
type UpdateInputDTO struct {
	ID          uuid.UUID
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}

type UpdateOutputDTO = GetOutputDTO

// DELETE
type DeleteInputDTO struct {
	ID uuid.UUID
}

// COUNT

type TotalInputDTO struct {
	StartDate   time.Time
	EndDate     time.Time
	UserID      *uuid.UUID
	ServiceName *string
}

type TotalOutputDTO struct {
	Total int64 `json:"total"`
}
