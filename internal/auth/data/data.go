package data

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Username  string
	Password  string
	CreatedAt string
}

func GetUser(conn *pgxpool.Pool, username string) (*User, error) {
	var user User
	err := conn.QueryRow(
		context.Background(),
		"select username, password, created_at::text from aether.user where username=$1",
		username,
	).Scan(&user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	return &user, nil
}
