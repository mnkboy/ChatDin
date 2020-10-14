package migratefunctions

import (
	"golangGraphQL/connection"

	usermodels "golangGraphQL/models/userModels"
	"golangGraphQL/settings"
)

// type UsuarioModelSqlserver struct {
// 	IdUsuario      int
// 	Codigo_Alterno string
// 	NickName       string
// 	Password       string
// 	RolId          int
// 	Estado         bool
// 	ImagenPerfil   string
// }

func ConsultaUsuarios() []usermodels.Usuarioserver {

	//Pedimos una conexion a la base de datos POSTGRES
	dbSqlServer := connection.OpenSQLSERVERConnection(settings.SqlServer)
	dbPostgres := connection.OpenConnection(settings.PostgresDB)

	//Por defecto siempre la cerramos
	defer dbSqlServer.Close()
	defer dbPostgres.Close()

	txServer := dbSqlServer.Begin()
	defer func() {
		if r := recover(); r != nil {
			txServer.Rollback()
		}
	}()

	// if err := txServer.Error; err != nil {
	// 	return err
	// }

	txPostgres := dbPostgres.Begin()

	defer func() {
		if r := recover(); r != nil {
			txPostgres.Rollback()
		}
	}()

	// if err2 := txPostgres.Error; err2 != nil {
	// 	return err2
	// }

	//Definimos una variable de la estructura de usuario

	usuarios := []usermodels.Usuarioserver{}

	//user := usermodels.UsuarioModel{}
	//consultamos los registros de la base de datos SQL SERVER
	//dbSqlServer.Find(&usuarios)
	txServer.Where("estadoSinc = ?", false).Find(&usuarios)

	for _, u := range usuarios {
		//creamos obj del struct de usaurio de postgres
		userP := usermodels.UsuarioModel{}
		//userP.IDUsuario = uuid.New()
		userP.CodigoAlterno = u.Codigo_Alterno
		userP.NickName = u.NickName
		userP.Password = u.Password
		userP.Estado = u.Estado
		userP.ImagenPerfil = u.ImagenPerfil
		result := txPostgres.Create(&userP).Error

		if result != nil {
			txPostgres.Rollback()
			txServer.Rollback()
		}

		result2 := txServer.Model(&u).Where("codigo_Alterno = ?", u.Codigo_Alterno).Update("estadoSinc", true).Error
		if result2 != nil {
			txPostgres.Rollback()
			txServer.Rollback()
		}

		//fmt.Println(result)

	}

	txPostgres.Commit()
	txServer.Commit()

	return usuarios

}

// func LoadData(db *gorm.DB) {
// 	//Pedimos una conexion a la base de datos POSTGRES
// 	dbSqlServer := connection.OpenConnection(settings.SqlServer)

// 	//Por defecto siempre la cerramos
// 	defer dbSqlServer.Close()

// 	//Definimos una variable de la estructura de usuario

// 	Allusers := usermodels.UsuarioModelSqlserver{}
// 	//user := usermodels.UsuarioModel{}
// 	//consultamos los registros de la base de datos SQL SERVER
// 	db.Find(&Allusers)
// 	/*
// 		for i, u := range Allusers {
// 			user.IDUsuario = u.IdUsuario
// 			user.CodigoAlterno = u.Codigo_Alterno
// 			user.NickName = u.NickName
// 			user.Password = u.Password
// 			user.Estado = u.Estado
// 			user.ImagenPerfil = u.ImagenPerfil

// 			//db.Create(&user)

// 		}
// 	*/

// 	//Definimos los parametros del usuario
// 	/*user := usermodels.UsuarioModel{
// 		CodigoAlterno: "123456",
// 		NickName:      "mnkboy",
// 		Password:      "123456",
// 		Estado:        true,
// 		ImagenPerfil:  "https://i.ytimg.com/vi/WezHBRv6Y4I/maxresdefault.jpg",
// 	}
// 	*/

// 	//Definimos los parametros del usuario
// 	/*user = usermodels.UsuarioModel{
// 		CodigoAlterno: "987654",
// 		NickName:      "whitemandrill",
// 		Password:      "159357",
// 		Estado:        false,
// 		ImagenPerfil:  "https://i.pinimg.com/originals/2a/44/9e/2a449ec03380a6cf6a9b4bf7d8ac3993.jpg",
// 	}

// 	//Insertar Usuario
// 	db.Create(&user)*/

// }
