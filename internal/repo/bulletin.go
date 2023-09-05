package repo

import (
	"database/sql"
	"fmt"
	"github.com/punkestu/buletin-go/internal/domain"
	"strings"
)

type Bulletin struct {
	conn *sql.DB
}

func NewBulletin(conn *sql.DB) *Bulletin {
	return &Bulletin{conn: conn}
}

func (b *Bulletin) Load(option *domain.BulletinOpt) ([]domain.Bulletin, error) {
	q := `SELECT * FROM bulletin_go.bulletin`
	if option != nil {
		var where []string
		if option.ID.Valid {
			where = append(where, fmt.Sprintf(`id=%#v`, option.ID.Int32))
		}
		if option.CreatorID.Valid {
			where = append(where, fmt.Sprintf(`creator_id=%#v`, option.CreatorID.Int32))
		}
		if len(where) > 0 {
			q += ` WHERE ` + strings.Join(where, " AND ")
		}
	}
	var result []domain.Bulletin
	dataRow, err := b.conn.Query(q)
	if err != nil {
		return nil, err
	}
	for dataRow.Next() {
		var buffer domain.Bulletin
		if err := dataRow.Scan(&buffer.ID, &buffer.Head, &buffer.Description, &buffer.CreatorID); err != nil {
			return nil, err
		}
		result = append(result, buffer)
	}
	return result, nil
}
func (b *Bulletin) Save(bulletin domain.Bulletin) (int32, error) {
	result, err := b.conn.Exec(`INSERT INTO bulletin_go.bulletin(head, description, creator_id) VALUES (?,?,?)`, bulletin.Head, bulletin.Description, bulletin.CreatorID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return int32(id), err
}
func (b *Bulletin) Delete(ID int32) error {
	_, err := b.conn.Exec("DELETE FROM bulletin_go.bulletin WHERE id = ?", ID)
	if err != nil {
		return err
	}
	return nil
}
func (b *Bulletin) LoadOne(option *domain.BulletinOpt) (*domain.Bulletin, error) {
	q := `SELECT * FROM bulletin_go.bulletin`
	if option != nil {
		var where []string
		if option.ID.Valid {
			where = append(where, fmt.Sprintf(`id=%#v`, option.ID.Int32))
		}
		if option.CreatorID.Valid {
			where = append(where, fmt.Sprintf(`creator_id=%#v`, option.CreatorID.Int32))
		}
		if len(where) > 0 {
			q += ` WHERE ` + strings.Join(where, " AND ")
		}
	}
	var buffer domain.Bulletin
	err := b.conn.QueryRow(q).Scan(&buffer.ID, &buffer.Head, &buffer.Description, &buffer.CreatorID)
	if err != nil {
		return nil, err
	}
	return &buffer, nil
}
