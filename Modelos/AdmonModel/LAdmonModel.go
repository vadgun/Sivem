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
