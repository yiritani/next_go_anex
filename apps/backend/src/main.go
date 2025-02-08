package main

import (
	"backend/src/cmd"
)

func main() {
	databasePath := "./db.sqlite"
	conn := sqlc.Connect(databasePath)

	defer conn.Close()

	queries := sqlc.New(conn)

	cmd.NewServer(queries)
}
