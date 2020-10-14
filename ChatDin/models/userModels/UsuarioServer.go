package usermodels

type Usuarioserver struct {
	UsuarioId      int64  `gorm:"column:usuarioId"`
	Codigo_Alterno string `gorm:"column:codigo_Alterno"`
	NickName       string `gorm:"column:nickName"`
	Password       string `gorm:"column:password"`
	RolId          int32  `gorm:"column:rolId"`
	Estado         bool   `gorm:"column:estado"`
	ImagenPerfil   string `gorm:"column:imagenPerfil"`
	EstadoSinc     bool   `gorm:"column:estadoSinc"`
}

//Cambiamos el nombre de la tabla a singular
func (Usuario Usuarioserver) TableName() string {
	return "usuarios"
}
