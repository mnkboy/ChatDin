package resolver

import (
	usermodels "chatDin/models/userModels"
)

//============= UsuarioResolvers =============

//UsuarioModelResolver :
type UsuarioModelResolver struct {
	Usuario usermodels.UsuarioModel
}

//IDUsuario :
func (r *UsuarioModelResolver) IDUsuario() string {
	return r.Usuario.IDUsuario
}

//CodigoAlterno :
func (r *UsuarioModelResolver) CodigoAlterno() *string {
	return &r.Usuario.CodigoAlterno
}

//NickName :
func (r *UsuarioModelResolver) NickName() *string {
	return &r.Usuario.NickName
}

//Password :
func (r *UsuarioModelResolver) Password() *string {
	return &r.Usuario.Password
}

//Estado :
func (r *UsuarioModelResolver) Estado() *bool {
	return &r.Usuario.Estado
}

//ImagenPerfil :
func (r *UsuarioModelResolver) ImagenPerfil() *string {
	return &r.Usuario.ImagenPerfil
}

//============= UsuarioResolvers =============
