package models

import "github.com/jinzhu/gorm"

//modelo contact para contactos

type Contact struct{
	gorm.Model

	Nombre		string	`json:"nombre"`
	Edad		string `json:"edad"`
	Telefono	string `json:telefono" gorm:"size:20"`
	Direccion	string `json:"direccion"`
	Email		string `json:"email"`
	Descripcion string `json:"descripcion" gorm:"type:TEXT"`
}