package main

import (
	"database/sql"
	"encoding/json"
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

	b, err := s.Create(service.BulletinCreate{
		Head: "hello",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	bm, err := json.Marshal(*b)
	log.Println(string(bm))
	err = s.Delete(b.ID)
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
}
