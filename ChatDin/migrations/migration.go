package main

import (
	"chatDin/connection"
	usermodels "chatDin/models/userModels"
	"fmt"
)

func main() {

	//Pedimos una conexion a la base de datos POSTGRES
	db := connection.OpenConnection("ChatDinPostgres")

	//Por defecto siempre la cerramos
	defer db.Close()

	//BorrarTabla
	msg := db.DropTable(&usermodels.UserAccessModel{})
	//Crear Tabla
	msg = db.CreateTable(&usermodels.UserAccessModel{})
	//Insertar Usuario
	msg = db.Create(&usermodels.UserAccessModel{
		CodigoAlterno: "123456",
		NickName:      "mnkboy",
		Password:      "123456",
		Estado:        true,
		ImagenPerfil:  "https://i.ytimg.com/vi/WezHBRv6Y4I/maxresdefault.jpg",
	})

	fmt.Println(msg)
}
