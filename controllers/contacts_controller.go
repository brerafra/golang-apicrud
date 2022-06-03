package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/brerafra/golang-apicrud/models"
	"github.com/brerafra/golang-apicrud/utils"
	"github.com/gorilla/mux"
)

//GetContact obtiene un contacto por su ID
func GetContact(w http.ResponseWriter, r *http.Request){
	//struct vacio donde se guardan los datos
	contact :=models.Contact{}
	//Se obtiene  el parametro id de la url
	id := mux.Vars(r)["id"]
	//Conexion a la BD
	bd := utils.GetConnection()
	defer db.Close()
	//Consulta a la db - SELECT * FROM contacts Where id=?
	db.Find(&contact, id)
	//Se comprueba que exista registro
	if contact.ID>0{
		//Se codifican los datos a formato json
		j, _:=json.Marshal(contact)
		//Se envian los datos
		utils.SendResponse(w, http.StatusOK, j)
	}else{
		//Si no existe se envia un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

//GetContacts obitene todos los contactos
func GetContacts(w http.ResponseWriter, r *http.Request){
	//Slice donde se guardan los datos
	contacts := []models.Contact{}
	//Conexion a la BD
	db := utils.GetConnection()
	defer db.Close()
	//Consulta a la bd-SELECT * FROM contacts
	db.Find(&contacts)

	//Se codifican los datos a formato json
	j, _ :=json.Marshal(contacts)
	//Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
}

//func StoreContact guarda un nuevo contact

func StoreContact (w http.ResponseWriter, r *http.Request){
	//Estructura donde se guardan los datos del body
	contact := models.Contact{}
	//Conexion a la BD
	db := utils.GetConnection()
	defer db.Close()

	//Se decodifican los datos del body a la estructura contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err!=nil{
		//Si hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}

	//Se guardan los datos en la BD
	err = db.Create(&contact).Error
	if err!=nil{
		//Si hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}
	//Se codfica el nuevo registro y se devuelve
	j,_ :=json.Marshal(contact)
	utils.SendResponse(w, http.StatusCreated,j)
}

//UpdateContact modifica los datos de un contacto por su id

func UpdateContact(w http.ResponseWriter, r *http.Request){
	//estructuras donde se almacenaran los datos
	contactFind := models.Contact{}
	contactData := models.Contact{}

	//Se obtiene el parametrio id de la url
	id := mux.Vars(r)["id"]
	//Conexion a la BD 
	db := utils.GetConnection()
	defer db.Close()

	//se buscan los datos
	db.Find(&contactFind, id)
	if contactFind.ID >0{
		//Si existe el registro se decodifican los datos del Body
		err := json.NewDecoder(r.Body).Decode(&contactData)
		if err!=nil{
			//si hay algun error en los datos se devolvera un error 400
			utils.SendErr(w, http.StatusBadRequest)
			return
		}

		//se modifican los datos
		db.Model(&contactFind).Updates(contactData)
		//se codifica el registro modificado y se devuelve

		j, _ :=json.Marshal(contactFind)
		utils.SendResponse(w, http.StatusOK, j)
	}else{
		//Si no existe el registro expecificado se devuelve un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
}

//DeleteContact elimina un contacto por ID
func DeleteContact(w http.ResponseWriter, r *http.Request){
	//Estructura donde se guardara el registro buscado
	contact := models.Contact{}
	//Se obtiene el parametro id de la url
	id := mux.Vars(r)["id"]

	//conexion a la BD
	db:= utils.GetConnection()
	defer db.Close()
	
	//Se busca el contacto
	db.Find(&contact, id)
	if contact.ID >0{
		//Si existe el contacto, se borra y se envia el contenido vacio
		db.Delete(contact)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	}else{
		//Si no existe el registro especificado se devuelve un error 404
		utils.SendErr(w, http.StatusNotFound)
	}
	
}