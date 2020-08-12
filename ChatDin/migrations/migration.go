package main

import (
	"chatDin/connection"
	"chatDin/migrations/migratefunctions"
)

func main() {
	//Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection("ChatDinPostgres")

	//Tiramos las tablas
	migratefunctions.DropTables(db)

	//Creamos las tablas
	migratefunctions.CreateTables(db)

	//Cargamos las tablas
	migratefunctions.LoadData(db)
}
