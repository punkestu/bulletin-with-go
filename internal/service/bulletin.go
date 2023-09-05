package service

import (
	"database/sql"
	"github.com/punkestu/buletin-go/internal/domain"
	"github.com/punkestu/buletin-go/internal/repo"
)

type Bulletin struct {
	b *repo.Bulletin
}

type BulletinCreate struct {
	Head        sql.NullString `json:"head"`
	Description sql.NullString `json:"description"`
	CreatorID   sql.NullInt32  `json:"creator_id"`
}

func NewBulletin(b *repo.Bulletin) *Bulletin {
	return &Bulletin{b: b}
}

func (b *Bulletin) Create(bulletin BulletinCreate) (*domain.Bulletin, error) {
	id, err := b.b.Save(domain.Bulletin{
		Head:        bulletin.Head,
		Description: bulletin.Description,
		CreatorID:   bulletin.CreatorID,
	})
	if err != nil {
		return nil, err
	}
	return b.b.LoadOne(&domain.BulletinOpt{
		ID: sql.NullInt32{
			Int32: id,
			Valid: true,
		},
	})
}

func (b *Bulletin) GetAll() ([]domain.Bulletin, error) {
	return b.b.Load(nil)
}

func (b *Bulletin) Delete(ID int32) error {
	return b.b.Delete(ID)
}

func (b *Bulletin) GetByID(bulletinID int32) (*domain.Bulletin, error) {
	return b.b.LoadOne(&domain.BulletinOpt{
		ID: sql.NullInt32{
			Int32: bulletinID,
			Valid: true,
		},
	})
}

func (b *Bulletin) GetByCreator(creatorID int32) ([]domain.Bulletin, error) {
	return b.b.Load(&domain.BulletinOpt{
		CreatorID: sql.NullInt32{
			Int32: creatorID,
			Valid: true,
		},
	})
}
