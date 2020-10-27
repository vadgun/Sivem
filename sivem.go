package main

import (
	"github.com/kataras/iris/v12"
	admoncontroller "github.com/vadgun/Sivem/Controladores/AdmonController"
	indexcontroller "github.com/vadgun/Sivem/Controladores/IndexController"
	logincontroller "github.com/vadgun/Sivem/Controladores/LoginController"
)

func main() {
	app := iris.New()
	app.HandleDir("/Recursos", "./Recursos")
	app.RegisterView(iris.HTML("./Vistas", ".html").Reload(true))

	app.Get("/", logincontroller.Getlogin)
	app.Get("/login", logincontroller.Getlogin)
	app.Post("/login", logincontroller.Getlogin)

	app.Get("/logout", logincontroller.Getlogout)

	app.Get("/index", indexcontroller.Index)
	app.Post("/index", indexcontroller.Index)

	app.Get("/clientes", admoncontroller.Clientes)

	app.Get("/empleados", admoncontroller.Empleados)

	app.Get("/empresas", admoncontroller.Empresas)

	app.Get("/medios", admoncontroller.Medios)

	app.Get("/materiales", admoncontroller.Materiales)

	//E S P E C T A C U L A R E S
	app.Get("/espectaculares", admoncontroller.Espectaculares)
	app.Get("/nuevoespectacular", admoncontroller.EspectacularesNuevo)
	app.Post("/guardarespectacular", admoncontroller.GuardaEspectacular)
	app.Post("/imagenesespectacular", admoncontroller.ObtenerImagenes)
	app.Post("/eliminarespectacular", admoncontroller.EliminarEspectacular)
	app.Post("/generarfichadecliente", admoncontroller.GenerarFichaDeCliente)

	/*------------------------ V E N T A S -----------------------------------*/
	app.Get("/ventas", admoncontroller.Ventas)

	// V A L L A S
	app.Get("/vallas", admoncontroller.Vallas)
	app.Get("/nuevavalla", admoncontroller.NuevaValla)



 /* ---------------------- V A L L A S   M O V I L E S ---------------------------*/
	app.Get("/vallasmoviles", admoncontroller.VallasMoviles)
	app.Get("/nuevavallamovil", admoncontroller.NuevaVallaMovil)

	//A L T A S
	app.Post("/guardarclientes", admoncontroller.GuardaClientes)
	app.Post("/editandoclientes", admoncontroller.EditandoClientes)
	app.Post("/guardarempleados", admoncontroller.GuardaEmpleados)
	app.Post("/editandoempleados", admoncontroller.EditandoEmpleados)

	//E D I T A R
	app.Post("/obtenercliente", admoncontroller.ObtenerCliente)
	app.Post("/obtenerempleado", admoncontroller.ObtenerEmpleado)

	//E L I M I N A R
	app.Post("/eliminarcliente", admoncontroller.EliminarCliente)
	app.Post("/eliminarempleado", admoncontroller.EliminarEmpleado)

	//B U S Q U E D A S
	app.Post("/verificaespectacular", admoncontroller.VerificaEspectacular)

	//M A T E R I A L E S
	app.Post("/guardarmaterial", admoncontroller.GuardaMaterial)
	app.Post("/obtenermaterial", admoncontroller.ObtenerMaterial)
	app.Post("/editarmaterial", admoncontroller.EditandoMaterial)
	app.Post("/eliminarmaterial", admoncontroller.EliminarMaterial)

	app.Run(iris.Addr(":8080"))

}
