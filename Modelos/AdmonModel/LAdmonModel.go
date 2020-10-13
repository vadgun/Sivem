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
