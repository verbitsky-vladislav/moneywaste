package models

type UserCreate struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UserGet struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
