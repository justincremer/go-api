package storage

import (
	"log"

	"github.com/gocql/gocql"
)

type storage struct {
	db *gocql.Session
}

func SetupStorage() (*storage, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "go_db"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	if err != nil {
		return &storage{}, err
	}

	return &storage{db: session}, nil
}

func (s *storage) ListUsers() ([]string, error) {
	var user string
	var users []string

	iter := s.db.Query(`SELECT name FROM users`).Iter()
	for iter.Scan(&user) {
		users = append(users, user)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return users, nil
}
