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

<<<<<<< HEAD
	// Update from database
	updatedRows, err := UpdateEmployee("Jake", "Poland")
	if err != nil {
		log.Fatal("Error updating Employee: ", err.Error())
	}
	fmt.Printf("Updated %d row(s) successfully.\n", updatedRows)

	// Delete from database
	deletedRows, err := DeleteEmployee("Jake")
	if err != nil {
		log.Fatal("Error deleting Employee: ", err.Error())
	}
	fmt.Printf("Deleted %d row(s) successfully.\n", deletedRows)
	*/
}

// CreateUser inserts an usuarios record
func CreateUser(idUsuario string, nickname string, password string, estado bool, imagenPerfil string) (string, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		err = errors.New("CreateUser: db is null")
		return "error", err
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return "error", err
	}

	tsql := "INSERT INTO usuario (idUsuario, nickname, password, estado, imagenPerfil) VALUES (@idUsuario, @nickname, @password , @estado , @imagenPerfil);"

	stmt, err := db.Prepare(tsql)
	if err != nil {
		return "error", err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("idUsuario", idUsuario),
		sql.Named("nickname", nickname),
		sql.Named("password", password),
		sql.Named("estado", estado),
		sql.Named("imagenPerfil", imagenPerfil))

	var newidUsuario string

	err = row.Scan(&newidUsuario)
	if err != nil {
		return "error", err
	}

	return newidUsuario, nil
=======
>>>>>>> angelcasanovasprint1
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
