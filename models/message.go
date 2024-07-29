package models

type Message struct {
	Id      int    `json:"id" db:"id"`
	Content string `json:"content" db:"content" binding:"required"`
}
