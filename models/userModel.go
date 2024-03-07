package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
	Username string `json:"username"`
}

type UserLinks struct {
	Links []map[string]string `bson:"Links"`
}