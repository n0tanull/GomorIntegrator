package app

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"simpleauth/internal/controller"
	"simpleauth/pkg/melog"
)

type Application interface {
	Run()
}
type app struct {
	con   controller.Controllerer
	melog melog.Logger
}

func NewApp() *app {
	me := melog.New()
	return &app{
		con:   controller.NewController(&me),
		melog: me,
	}
}

func (a *app) Run() {
	r := fiber.New()
	a.melog.Debug("starting server")
	r.Get("/user", a.con.GetUser)
	r.Get("/users", a.con.GetUsers)
	r.Post("/user", a.con.AddUser)
	r.Post("/social", a.con.AddSocial)
	a.melog.Info("server started")
	log.Fatal(r.Listen(":8080"))
}
