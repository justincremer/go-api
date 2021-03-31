package storage

import (
	"log"

	"github.com/google/uuid"
	"github.com/justincremer/go-api/pkg/schemas"
)

func (storage *storage) ListUsers() ([]string, error) {
	var user string
	var users []string

	query := `SELECT username FROM users`
	iter := storage.db.Query(query).Iter()
	for iter.Scan(&user) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}

func (storage *storage) CreateUser(user schemas.User) (string, error) {
	id := uuid.New().String()

	query := `INSERT INTO users (user_id, username, name, email, phone) Values (?, ?, ?, ?, ?)`
	if err := storage.db.Query(query, id, user.Username, user.Name, user.Email, user.Phone).Exec(); err != nil {
		log.Println("Error creating writing user:", err)
		return "", err
	}

	return id, nil
}

func (storage *storage) UpdateUser(user schemas.User) error {
	query := `UPDATE users SET username=?, name=?, email=?, phone=? WHERE user_id=?`
	if err := storage.db.Query(query, user.Username, user.Name, user.Email, user.Phone, user.Id).Exec(); err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func (storage *storage) DeleteUser(user schemas.User) error {
	query := `DELETE FROM users WEHRE user_id=?`
	if err := storage.db.Query(query, user.Id).Exec(); err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	return nil
}
