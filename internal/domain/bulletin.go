package domain

import "database/sql"

type Bulletin struct {
	ID          int32          `json:"id"`
	Head        sql.NullString `json:"head"`
	Description sql.NullString `json:"description"`
	CreatorID   sql.NullInt32  `json:"creator_id"`
}

type BulletinOpt struct {
	ID        sql.NullInt32
	CreatorID sql.NullInt32
}
