package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Cart struct {
	bun.BaseModel `bun:"table:cart,alias:c"`

	Id int64 `json:"-" bun:"id,pk,autoincrement"`

	CartId uuid.UUID `json:"id,omitempty" bun:"type:uuid,default:uuid_generate_v4()"`

	UserId string `json:"userId,omitempty"`

	// Items []Item `json:"items,omitempty"`

	TotalQuantity int32 `json:"totalQuantity,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty" bun:"type:timestamp,default:now()"`
}
