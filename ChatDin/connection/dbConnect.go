package connection

import (
	"encoding/xml"
	"fmt"
	"golangGraphQL/models/settingsModels"
	"io/ioutil"
	"os"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//OpenConnection : abrimos conexion a la base de datos
func OpenConnection(database string) *gorm.DB {
	//Abrimos archivo
	xmlFile, err := os.Open("../settings/BDSettings.xml")

	//Verificamos que no existan errores
	if err != nil {
		panic(err)
	}

	fmt.Println("Archivo BdSettings.xml abierto exitosamente")

	//Por defecto siempre cerramos el archivo
	defer xmlFile.Close()

	//Leemos en un byte array el contenido del archivo
	dbSettingsByteArray, _ := ioutil.ReadAll(xmlFile)

	//Declaramos variables
	var DBs settingsModels.DBSettingsModel

	//volcamos el contenido del byte array en las estructuras
	xml.Unmarshal(dbSettingsByteArray, &DBs)

	//Creamos un item de tipo DB
	DBItem := &settingsModels.DBModel{}

	//Imprimimos las configuraciones
	for _, dbItem := range DBs.DataBase {
		if dbItem.Name == database {
			DBItem = &dbItem
			break
		}
	}

	//Creamos la cadena de conexion
	strcon := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%s",
		DBItem.Server,
		DBItem.Port,
		DBItem.User,
		DBItem.Database,
		DBItem.Password,
		DBItem.SslMode,
	)

	//Abrimos una conexion a la base de datos
	db, err := gorm.Open(DBItem.Engine, strcon)

	//Si el error es distinto a nil tiramos el error
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Conexion abierta exitosamente a base dedatos: " + DBItem.Name)
	}

	//Habilitamos el log de actividades
	db.LogMode(true)

	//Hacemos ping al servidor
	db.DB().Ping()

	//Creamos la extension para uuid
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	//Devolvemos la conexion abierta a la base de datos
	return db
}
