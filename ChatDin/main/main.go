package main

import (
	"chatDin/models/settingsModels"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Abrimos archivo
	xmlFile, err := os.Open("../settings/BDSettings.xml")

	//Verificamos que no existan errores
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully opened BdSettings.xml")

	//Por defecto siempre cerramos el archivo
	defer xmlFile.Close()

	//Leemos en un byte array el contenido del archivo
	dbSettingsByteArray, _ := ioutil.ReadAll(xmlFile)

	//Declaramos variables
	var DBs settingsModels.DBSettingsModel

	//volcamos el contenido del byte array en las estructuras
	xml.Unmarshal(dbSettingsByteArray, &DBs)

	//Imprimimos las configuraciones
	for _, dbItem := range DBs.DataBase {
		fmt.Println("==============================")
		fmt.Println("Name: " + dbItem.Name)
		fmt.Println("Engine: " + dbItem.Engine)
		fmt.Println("Server: " + dbItem.Server)
		fmt.Println("Port: " + dbItem.Port)
		fmt.Println("User: " + dbItem.User)
		fmt.Println("Password: " + dbItem.Password)
		fmt.Println("Database: " + dbItem.Database)
		fmt.Println("SslMode: " + dbItem.SslMode)
		fmt.Println("==============================")
	}

}
