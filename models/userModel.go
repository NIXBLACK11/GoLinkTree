package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Token string `json:"token"`
}

type UserLinks struct {
	Links []map[string]string `bson:"Links"`
}

type Link struct {
	Name string `json:"name"`
	URL string `json:"url"`
}