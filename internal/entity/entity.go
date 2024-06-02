package entity

type Client struct {
	Id       int    `json:"id"`
	Telegram string `json:"Telegram"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
