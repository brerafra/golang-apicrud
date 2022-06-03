package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/brerafra/golang-apicrud/routes"
	"github.com/brerafra/golang-apicrud/utils"
	"github.com/gorilla/mux"
)

func main(){
	//Migracion de la base de datos
	utils.MigrateDB()

	//router para el manejo de las rutas de la aplicacion
	r := mux.NewRouter()

	//Se agregan las rutas de contactos 
	routes.SetContactsRoutes(r)

	//Generacion  de unuevo servidor, especificamos el puerto y las rutas
	srv := http.Server{
		Addr: ":4000",
		Handler: r,
	}

	log.Println("Running on port 4000")
	//Se ejecuta el servidor
	log.Println(srv.ListenAndServe())
}