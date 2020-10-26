package admonmodel

import (
	"time"

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

//EmpleadoMongo Controla datos de los Empledos en el Sistema
type EmpleadoMongo struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Control     string        `bson:"Control"`
	Nombre      string        `bson:"Nombre"`
	ApellidoP   string        `bson:"ApellidoP"`
	ApellidoM   string        `bson:"ApellidoM"`
	Sexo        string        `bson:"Sexo"`
	Cargo       string        `bson:"Cargo"`
	Licencia    string        `bson:"Licencia"`
	Correo      string        `bson:"Correo"`
	Calle       string        `bson:"Calle"`
	Numero      string        `bson:"Numero"`
	Colonia     string        `bson:"Colonia"`
	Municipio   string        `bson:"Municipio"`
	Estado      string        `bson:"Estado"`
	Ciudad      string        `bson:"Ciudad"`
	CodigoP     string        `bson:"CodigoP"`
	Telefono    string        `bson:"Telefono"`
	Referencias string        `bson:"Referencias"`
}

//EspectacularMongo -> Controla la estructura de los espectaculares en la base de datos
type EspectacularMongo struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	NumControl       string        `bson:"NumControl"`
	CostoImpresion   float64       `bson:"CostoImpresion"`
	CostoInstalacion float64       `bson:"CostoInstalacion"`
	Calle            string        `bson:"Calle"`
	Numero           string        `bson:"Numero"`
	Colonia          string        `bson:"Colonia"`
	Localidad        string        `bson:"Localidad"`
	Municipio        string        `bson:"Municipio"`
	Estado           string        `bson:"Estado"`
	Latitud          string        `bson:"Latitud"`
	Longitud         string        `bson:"Longitud"`
	Referencias      string        `bson:"Referencias"`
	Ancho            float64       `bson:"Ancho"`
	Alto             float64       `bson:"Alto"`
	Material         MaterialMongo `bson:"Material"`
	Precio           float64       `bson:"Precio"`
	Observaciones    string        `bson:"Observaciones"`
	Status           string        `bson:"Status"`
	Acabados         string        `bson:"Acabados"`
	Propietario      string        `bson:"Propietario"`
	Celular          string        `bson:"Celular"`
	Telefono         string        `bson:"Telefono"`
	Imagenes         []string      `bson:"Imagenes"`
	ContratoInicio   time.Time     `bson:"ContratoInicio"`
	ContratoFin      time.Time     `bson:"ContratoFin"`
	Monto            float64       `bson:"Monto"`
	Folio            string        `bson:"Folio"`
	TipoPago         string        `bson:"TipoPago"`
	PeriodoPago      string        `bson:"PeriodoPago"`
	Disponible       bool          `bson:"Disponible"`
	Rentado          bool          `bson:"Rentado"`
}

//MaterialMongo -> Estructura de material para la base de datos
type MaterialMongo struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Nombre      string        `bson:"Nombre"`
	Precio      float64       `bson:"Precio"`
	Descripcion string        `bson:"Descripcion"`
}

//EmpresaMongo -> Estructura que controla la empresa en la base de datos
type EmpresaMongo struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Nombre string        `bson:"Nombre"`
	RFC    string        `bson:"RFC"`
}
