package domain

import (
	"database/sql"
)

type Bulletin struct {
	ID          int32  `json:"id"`
	Head        string `json:"head"`
	Description string `json:"description"`
	CreatorID   *int32 `json:"creator_id,omitempty"`
}

type BulletinOpt struct {
	ID        sql.NullInt32
	CreatorID sql.NullInt32
}
