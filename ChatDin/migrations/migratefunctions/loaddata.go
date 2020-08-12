package migratefunctions

import (
	usermodels "chatDin/models/userModels"

	"github.com/jinzhu/gorm"
)

func LoadData(db *gorm.DB) {
	//Definimos los parametros del usuario
	user := usermodels.UsuarioModel{
		CodigoAlterno: "123456",
		NickName:      "mnkboy",
		Password:      "123456",
		Estado:        true,
		ImagenPerfil:  "https://i.ytimg.com/vi/WezHBRv6Y4I/maxresdefault.jpg",
	}

	//Insertar Usuario
	db.Create(&user)

	//Definimos los parametros del usuario
	user = usermodels.UsuarioModel{
		CodigoAlterno: "987654",
		NickName:      "whitemandrill",
		Password:      "159357",
		Estado:        false,
		ImagenPerfil:  "https://i.pinimg.com/originals/2a/44/9e/2a449ec03380a6cf6a9b4bf7d8ac3993.jpg",
	}

	//Insertar Usuario
	db.Create(&user)

}
