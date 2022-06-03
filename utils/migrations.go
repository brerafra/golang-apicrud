package utils

import(
	"fmt"
	"golang-apicrud/models"
)

//MigrateDB migra la base de datos

func MigrateDB(){
	db := GetConnection()
	defer db.Close()
	fmt.Println("Migration models...")

	//Automigrate se encarga de mgirar la base de datos si no se ha migrado,
	//y lo hace apartir del modelo

	db.AutoMigrate(&models.Contact{})
}