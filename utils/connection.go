package utils

import(
	"log"
	"github.com/jinzhu/gorm"
	//Mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

//Get conection obtiene una conexión a la base de datos

func GetConnection() *gorm.DB{
	db, err := gorm.Open("mysql","root:demosprosa#1@/pruebas_go?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Fatal(err)
	}
	return db
}