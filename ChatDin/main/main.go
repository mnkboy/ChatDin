package main

import (
	"chatDin/connection"
)

func main() {
	DB := connection.OpenConnection("ChatDinPostgres")

	DB = DB
}
