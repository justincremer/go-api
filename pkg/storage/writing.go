package storage

import (
	"log"

	"github.com/google/uuid"
	writing "github.com/justincremer/go-api/pkg/write"
)

func (storage *storage) CreateUser(user writing.User) (string, error) {
	id := uuid.New().String()

	if err := storage.db.Query(`INSERT INTO users (user_id, username, name, email, phone) Values (?, ?, ?, ?, ?)`,
		id, user.Username, user.Name, user.Email, user.Phone).Exec(); err != nil {
		log.Println("Error writing user to DB: ", err)
		return "", err
	}

	return id, nil
}
