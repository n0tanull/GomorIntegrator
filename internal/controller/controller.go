package controller

import (
	"github.com/gofiber/fiber/v3"
	"simpleauth/internal/service"
)

type Controllerer interface {
	AddUser(con fiber.Ctx) error
	AddSocial(con fiber.Ctx) error
	GetUser(con fiber.Ctx) error
	GetUsers(con fiber.Ctx) error
}

type controller struct {
	service.Service
}

func NewController(repo service.Service) Controllerer {
	return &controller{
		Service: service.NewService(repo),
	}
}

func (c *controller) AddUser(con fiber.Ctx) error {
	email := con.Query("email")
	password := con.Query("password")
	if email == "" || password == "" {
		return fiber.ErrBadRequest
	}
	res, err := c.Service.AddUser(email, password, con.UserContext())
	if err != nil {
		return err
	}
	return con.SendString(res)
}

func (c *controller) AddSocial(con fiber.Ctx) error {
	telegram := con.Query("telegram")
	if telegram == "" {
		return fiber.ErrBadRequest
	}
	res, err := c.Service.AddSocial(telegram, con.UserContext())
	if err != nil {
		return err
	}
	return con.SendString(res)
}

func (c *controller) GetUser(con fiber.Ctx) error {
	id := con.Query("id")
	if id == "" {
		return fiber.ErrBadRequest
	}

	res, err := c.Service.GetUser(id, con.UserContext())
	if err != nil {
		return err
	}
	return con.SendString(res)
}

func (c *controller) GetUsers(con fiber.Ctx) error {
	res, err := c.Service.GetUsers(con.UserContext())
	if err != nil {
		return err
	}
	return con.JSON(res)
}
