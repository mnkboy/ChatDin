package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "SA"
var password = "Server@2020"
var database = "testdb"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	//Read usuarios
	count, err := ReadUsuarios()
	if err != nil {
		log.Fatal("Error reading Usuarios: ", err.Error())
	}
	fmt.Printf("Read %d row(s) successfully.\n", count)

}

// ReadUsuarios reads all employee records
func ReadUsuarios() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	tsql := fmt.Sprintf("SELECT idUsuario, nickname, password, estado, imagenPerfil FROM usuarios;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return -1, err
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	for rows.Next() {
		var idUsuario, nickname, password, imagenPerfil string
		var estado bool

		// Get values from row.
		err := rows.Scan(&idUsuario, &nickname, &password, &estado, &imagenPerfil)
		if err != nil {
			return -1, err
		}

		fmt.Printf("ID: %s, nickname: %s, Password: %s, Estado: %t, UbicacionFoto: %s \n", idUsuario, nickname, password, estado, imagenPerfil)
		count++
	}

	return count, nil
}
