package usermodels

//UsuarioModel :
type UsuarioModel struct {
	IDUsuario     string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CodigoAlterno string `gorm:"type:varchar(30)";`
	NickName      string `gorm:"type:varchar(60)";`
	Password      string `gorm:"type:varchar(200)";`
	Estado        bool   `gorm:"type:bool";`
	ImagenPerfil  string `gorm:"type:varchar(300)";`
}

//Cambiamos el nombre de la tabla a singular
func (Usuario UsuarioModel) TableName() string {
	return "usuario"
}
