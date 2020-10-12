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

	app.Get("/espectaculares", admoncontroller.Espectaculares)

	app.Get("/ventas", admoncontroller.Ventas)

	app.Get("/vallas", admoncontroller.Vallas)

	app.Get("/vallasmoviles", admoncontroller.VallasMoviles)

	//A L T A S
	app.Post("/guardarclientes", admoncontroller.GuardaClientes)
	app.Post("/editandoclientes", admoncontroller.EditandoClientes)

	//E D I T A R
	app.Post("/obtenercliente", admoncontroller.ObtenerCliente)

	//E L I M I N A R
	app.Post("/eliminarcliente", admoncontroller.EliminarCliente)

	app.Run(iris.Addr(":8080"))

}
