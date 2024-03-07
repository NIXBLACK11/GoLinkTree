package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
	Username string `json:"username"`
}

type LinkDetails struct {
	Name string `json:"name"`
	Link string `json:"link"`
}