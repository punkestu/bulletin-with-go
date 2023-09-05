package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/punkestu/buletin-go/internal/repo"
	"github.com/punkestu/buletin-go/internal/service"
	"log"
)

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bulletin_go")
	if err != nil {
		log.Fatalln(err)
		return
	}
	r := repo.NewBulletin(conn)
	s := service.NewBulletin(r)
	created, err := s.Create(service.BulletinCreate{
		Head: sql.NullString{
			String: "Hello",
			Valid:  true,
		},
		Description: sql.NullString{
			String: "ini hello",
			Valid:  true,
		},
		CreatorID: sql.NullInt32{
			Int32: 100,
			Valid: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	bs, err := s.GetAll()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(bs)
	bs, err = s.GetByCreator(100)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(bs)
	b, err := s.GetByID(created.ID)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(*b)
	err = s.Delete(created.ID)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
