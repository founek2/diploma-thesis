package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Item struct {
	bun.BaseModel `bun:"table:item,alias:i"`

	Id int64 `json:"-" bun:"id,pk,autoincrement"`

	ItemId uuid.UUID `json:"id,omitempty" bun:"type:uuid,default:uuid_generate_v4()"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	Quantity int32 `json:"quantity,omitempty"`

	Price float64 `json:"price,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty" bun:"type:timestamp,default:now()"`
}
