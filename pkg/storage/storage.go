package storage

import (
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
