package admonmodel

import (
	"gopkg.in/mgo.v2/bson"
)

//ClienteMongo Controla datos del Cliente en el sistema
type ClienteMongo struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Nombre        string        `bson:"Nombre"`
	ApellidoP     string        `bson:"ApellidoP"`
	ApellidoM     string        `bson:"ApellidoM"`
	Sexo          string        `bson:"Sexo"`
	Empresa       string        `bson:"Empresa"`
	Puesto        string        `bson:"Puesto"`
	Calle         string        `bson:"Calle"`
	Numero        string        `bson:"Numero"`
	Colonia       string        `bson:"Colonia"`
	Municipio     string        `bson:"Municipio"`
	Estado        string        `bson:"Estado"`
	Ciudad        string        `bson:"Ciudad"`
	CodigoP       string        `bson:"CodigoP"`
	Correo        string        `bson:"Correo"`
	CorreoEMP     string        `bson:"CorreoEMP"`
	Telefono      string        `bson:"Telefono"`
	TelefonoA     string        `bson:"TelefonoA"`
	SitioWeb      string        `bson:"SitioWeb"`
	Observaciones string        `bson:"Observaciones"`
}
