package repository

import (
	"MessageProcessing/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) Create(message models.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (content)"+
		"VALUES ($1) RETURNING id", messageTable)
	row := r.db.QueryRow(query, message.Content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MessagePostgres) GetCurMessages() ([]models.Message, error) {
	var curMessages []models.Message
	curQuery := fmt.Sprintf("SELECT id, content FROM %s WHERE status_id=1", messageTable)
	err := r.db.Select(&curMessages, curQuery)

	return curMessages, err
}

func (r *MessagePostgres) GetCompMessages() ([]models.Message, error) {
	var compMessages []models.Message

	compQuery := fmt.Sprintf("SELECT id, content FROM %s WHERE status_id=2", messageTable)
	err := r.db.Select(&compMessages, compQuery)

	return compMessages, err
}

func (r *MessagePostgres) UpdateStatus(messageId int) error {
	query := fmt.Sprintf("UPDATE %s SET status_id = 2 WHERE id = $1", messageTable)
	_, err := r.db.Exec(query, messageId)

	return err
}
