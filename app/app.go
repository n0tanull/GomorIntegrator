package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"simpleauth/internal/controller"
	"simpleauth/internal/repo/user"
	"simpleauth/internal/service"
)

type Application interface {
	Run()
}
type app struct {
	con controller.Controllerer
}

func NewApp() *app {
	con := controller.NewController(service.NewService(user.NewRepo()))
	return &app{
		con: con,
	}
}

func (a *app) Run() {
	r := fiber.New()
	r.Get("/user", a.con.GetUser)
	r.Get("/users", a.con.GetUsers)
	r.Post("/user", a.con.AddUser)
	r.Post("/social", a.con.AddSocial)
	log.Fatal(r.Listen(":8080"))
}
