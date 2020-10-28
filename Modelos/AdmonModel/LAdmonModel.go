package admonmodel

import (
	"fmt"
	"log"

	conexiones "github.com/vadgun/Sivem/Conexiones"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ExtraeClientes -> Extrae los clientes y los regresa todos
func ExtraeClientes() []ClienteMongo {

	var clientes []ClienteMongo

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_CL)
	err1 := c.Find(bson.M{}).All(&clientes)
	if err1 != nil {
		fmt.Println(err1)
	}

	return clientes
}

//ExtraeEmpleados -> Devuelve la coleccion de Empleados a la vista
func ExtraeEmpleados() []EmpleadoMongo {
	var empleados []EmpleadoMongo

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_EM)
	err1 := c.Find(bson.M{}).All(&empleados)
	if err1 != nil {
		fmt.Println(err1)
	}

	return empleados
}

//ExtraeEspectaculares -> Devuelve la coleccion de Espectaculares a la vista
func ExtraeEspectaculares() []EspectacularMongo {
	var espectaculares []EspectacularMongo

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.Find(bson.M{}).All(&espectaculares)
	if err1 != nil {
		fmt.Println(err1)
	}

	return espectaculares

}

//ExtraeMateriales Regresa la coleccion completa de materiales a la vista
func ExtraeMateriales() []MaterialMongo {
	var materiales []MaterialMongo

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MAT)
	err1 := c.Find(bson.M{}).All(&materiales)
	if err1 != nil {
		fmt.Println(err1)
	}

	return materiales

}

//ObtenerEspectacularPorID -> Regresa la informacion del espectacular por ID
func ObtenerEspectacularPorID(id string) EspectacularMongo {
	objid := bson.ObjectIdHex(id)
	var espectacular EspectacularMongo
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.FindId(objid).One(&espectacular)
	if err1 != nil {
		fmt.Println(err1)
	}
	return espectacular
}

//GuardarClienteMongo -> Guarda Clientes en la Base de datos de Mongo DB
func GuardarClienteMongo(cliente ClienteMongo) bool {

	var guardado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_CL)
	err1 := c.Insert(cliente)
	guardado = true
	if err1 != nil {
		guardado = false
		fmt.Println(err1)
	}

	return guardado

}

//GuardaEmpleadoMongo -> Guarda los datos del Empleado en la Base de datos
func GuardaEmpleadoMongo(empleado EmpleadoMongo) bool {

	var guardado bool
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_EM)
	err1 := c.Insert(empleado)
	guardado = true
	if err1 != nil {
		guardado = false
		fmt.Println(err1)
	}

	return guardado

}

//EditarClienteMongo -> Guarda los cambios en la Base de datos
func EditarClienteMongo(id string, cliente ClienteMongo) bool {

	objid := bson.ObjectIdHex(id)
	var editado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_CL)
	err1 := c.UpdateId(objid, cliente)
	editado = true
	if err1 != nil {
		editado = false
		fmt.Println(err1)
	}

	return editado

}

//EditarEmpleadoMongo -> Edita el empleado en la base de datos
func EditarEmpleadoMongo(id string, empleado EmpleadoMongo) bool {
	objid := bson.ObjectIdHex(id)
	var editado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_EM)
	err1 := c.UpdateId(objid, empleado)
	editado = true
	if err1 != nil {
		editado = false
		fmt.Println(err1)
	}

	return editado

}

//ExtraeClientePorId -> Regresa los datos del Cliente para ser editado
func ExtraeClientePorId(id string) ClienteMongo {

	var cliente ClienteMongo

	objid := bson.ObjectIdHex(id)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_CL)
	err1 := c.FindId(objid).One(&cliente)
	if err1 != nil {
		fmt.Println(err1)
	}

	return cliente

}

//ExtraeEmpleadoPorId -> Extrae Empleado por medio de la id
func ExtraeEmpleadoPorId(id string) EmpleadoMongo {
	var empleado EmpleadoMongo

	objid := bson.ObjectIdHex(id)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_EM)
	err1 := c.FindId(objid).One(&empleado)
	if err1 != nil {
		fmt.Println(err1)
	}

	return empleado
}

//EliminarEspectacularMongo -> Elimina el espectacular de la base de datos
func EliminarEspectacularMongo(id string) {
	objid := bson.ObjectIdHex(id)
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.RemoveId(objid)
	if err1 != nil {
		fmt.Println(err1)
	}
}

//EliminarClienteMongo -> Elimina el Cliente de la base de datos
func EliminarClienteMongo(id string) bool {
	objid := bson.ObjectIdHex(id)
	var eliminado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_CL)
	err1 := c.RemoveId(objid)
	eliminado = true
	if err1 != nil {
		eliminado = false
		fmt.Println(err1)
	}

	return eliminado
}

//EliminarEmpleadoMongo -> Elimina un empleado de la base de datos
func EliminarEmpleadoMongo(id string) bool {
	objid := bson.ObjectIdHex(id)
	var eliminado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_EM)
	err1 := c.RemoveId(objid)
	eliminado = true
	if err1 != nil {
		eliminado = false
		fmt.Println(err1)
	}

	return eliminado
}

//GuardarEspectacularMongo -> Guarda el espectacular y sus imagenes en la base de datos
func GuardarEspectacularMongo(espectacular EspectacularMongo) bool {
	var guardado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.Insert(espectacular)
	guardado = true
	if err1 != nil {
		guardado = false
		fmt.Println(err1)
	}

	return guardado
}

//VerificaEspectacularMongo -> Verifica el numero de control para el alta de espectaculares
func VerificaEspectacularMongo(numcontrol string) bool {

	var espectacular EspectacularMongo

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.Find(bson.M{"NumControl": numcontrol}).One(&espectacular)
	if err1 != nil {
		fmt.Println(err1)
	}

	if espectacular.NumControl == "" {
		return false
	} else {
		return true
	}

}

//GuardarMaterialMongo -> Guarda el material en la base de datos de mongodb
func GuardarMaterialMongo(material MaterialMongo) bool {
	var guardado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MAT)
	err1 := c.Insert(material)
	guardado = true
	if err1 != nil {
		guardado = false
		fmt.Println(err1)
	}

	return guardado
}

//ExtraeMaterialPorId -> Extrae Material por medio de la id
func ExtraeMaterialPorId(id string) MaterialMongo {
	var material MaterialMongo

	objid := bson.ObjectIdHex(id)

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MAT)
	err1 := c.FindId(objid).One(&material)
	if err1 != nil {
		fmt.Println(err1)
	}

	return material
}

//EditarMaterialMongo -> Guarda los cambios en la Base de datos
func EditarMaterialMongo(id string, material MaterialMongo) bool {

	objid := bson.ObjectIdHex(id)
	var editado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MAT)
	err1 := c.UpdateId(objid, material)
	editado = true
	if err1 != nil {
		editado = false
		fmt.Println(err1)
	}

	return editado

}

//EliminarMaterialMongo -> Elimina el material de la base de datos
func EliminarMaterialMongo(id string) bool {
	objid := bson.ObjectIdHex(id)
	var eliminado bool

	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_MAT)
	err1 := c.RemoveId(objid)
	eliminado = true
	if err1 != nil {
		eliminado = false
		fmt.Println(err1)
	}

	return eliminado
}

//ExtraeEspectacularesDisponibles -> Extrae los Espectaculares unicamente disponibles
func ExtraeEspectacularesDisponibles() []EspectacularMongo {
	var espectaculares []EspectacularMongo
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.Find(bson.M{"Disponible": true}).All(&espectaculares)
	if err1 != nil {
		fmt.Println(err1)
	}

	return espectaculares

}

//ExtraeEspectacularesRentados -> Extrae los Espectaculares unicamente rentados
func ExtraeEspectacularesRentados() []EspectacularMongo {
	var espectaculares []EspectacularMongo
	session, err := mgo.Dial(conexiones.MONGO_SERVER)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	c := session.DB(conexiones.MONGO_DB).C(conexiones.MONGO_DB_ES)
	err1 := c.Find(bson.M{"Disponible": true, "Rentado": true}).All(&espectaculares)
	if err1 != nil {
		fmt.Println(err1)
	}

	return espectaculares

}
