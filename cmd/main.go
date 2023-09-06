package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/punkestu/buletin-go/internal/repo"
	"github.com/punkestu/buletin-go/internal/service"
	"log"
	"net/http"
)

func main() {
	conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bulletin_go")
	if err != nil {
		log.Fatalln(err)
		return
	}
	r := repo.NewBulletin(conn)
	s := service.NewBulletin(r)
	c := cors.New()
	app := fiber.New()

	bulletin := app.Group("/bulletin")

	bulletin.Use(c)

	bulletin.Post("/", func(ctx *fiber.Ctx) error {
		var r struct {
			Head        string `json:"head"`
			Description string `json:"description"`
		}
		if err := ctx.BodyParser(&r); err != nil {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		bulletin, err := s.Create(service.BulletinCreate{
			Head:        r.Head,
			Description: r.Description,
		})
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}
		return ctx.JSON(bulletin)
	})

	bulletin.Get("/", func(ctx *fiber.Ctx) error {
		bulletins, err := s.GetAll()
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}
		return ctx.JSON(bulletins)
	})

	bulletin.Get("/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		bulletin, err := s.GetByID(int32(id))
		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}
		return ctx.JSON(bulletin)
	})

	bulletin.Delete("/:id", func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return ctx.SendStatus(http.StatusBadRequest)
		}
		if err := s.Delete(int32(id)); err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}
		return ctx.SendStatus(http.StatusOK)
	})

	err = app.Listen(":8000")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
