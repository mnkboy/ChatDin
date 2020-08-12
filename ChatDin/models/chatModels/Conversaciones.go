package chatmodels

import (
	usermodels "chatDin/models/userModels"
)

//ConversacionesModel :
type ConversacionesModel struct {
	IDConversacion          string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	IDUsuarioEmisor         string `gorm:"type:varchar(30)";`
	UsuarioEmisor           usermodels.UsuarioModel
	IDUsuarioReceptor       string `gorm:"type:varchar(60)";`
	UsuarioReceptor         usermodels.UsuarioModel
	FechaDeCreacion         string `gorm:"type:varchar(200)";`
	EstatusConversacion     bool   `gorm:"type:bool";`
	SilenciarNotificaciones bool   `gorm:"type:varchar(300)";`
	Leido                   bool   `gorm:"type:varchar(300)";`
	Archivar                bool   `gorm:"type:varchar(300)";`
	Tipo                    string `gorm:"type:varchar(30)";`
}

//Cambiamos el nombre de la tabla a singular
func (Conversaciones ConversacionesModel) TableName() string {
	return "conversaciones"
}
