package resolver

import (
	"fmt"
	usermodels "chatDin/models/userModels"

	"github.com/jinzhu/gorm"
)

//Resolver :
type Resolver struct {
	DB *gorm.DB
}

//Hello :
func (r *Resolver) Hello(args struct{ ID string }) (string, error) {
	arrayModel := []usermodels.UsuarioModel{}

	//Obtenemos los where de acuerdo a los campos del modelo
	db := r.DB.Where("id_usuario = ?", args.ID)

	//Realizamos la consulta
	db.Find(&arrayModel)

	c := len(arrayModel)
	fmt.Println(c)

	if len(arrayModel) > 0 {
		return fmt.Sprintf("Hello, %v", arrayModel[0].NickName), nil
	}
	return "", nil

}

//Usuarios :
func (r *Resolver) Usuarios(args struct{ IDs []string }) ([]*UsuarioModelResolver, error) {

	arrayModel := &[]usermodels.UsuarioModel{}

	//Obtenemos los where de acuerdo a los campos del modelo
	db := r.DB.Where("id_usuario in (?)", args.IDs)

	//Realizamos la consulta
	db.Find(&arrayModel)

	c := len(*arrayModel)

	resolver := []*UsuarioModelResolver{}
	if c > 0 {
		for _, item := range *arrayModel {

			resolver = append(resolver,
				&UsuarioModelResolver{item},
			)
		}
		//Devolvemos array lleno
		return resolver, nil
	}
	//Devolvemos array vacio
	return resolver, nil

}
