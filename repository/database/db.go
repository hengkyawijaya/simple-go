package repository_database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hengkyawijaya/simple-go/repository/config"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(ctx context.Context, cfg *repository_config.Config) *Database {
	psqlClient := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Repository.Database.Host,
		cfg.Repository.Database.Port,
		cfg.Repository.Database.Username,
		cfg.Repository.Database.Password,
		cfg.Repository.Database.Name)

	db, err := sql.Open("postgres", psqlClient)
	if err != nil {
		panic(err)
	}

	return &Database{
		Conn: db,
	}
}

func (d *Database) GetUser(ctx context.Context, id int) (string, error) {
	var name string
	err := d.Conn.QueryRow("SELECT name FROM users WHERE user_id = $1", id).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
