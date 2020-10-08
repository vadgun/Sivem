package main

import (
	"github.com/kataras/iris/v12"
	logincontroller "github.com/vadgun/Sivem/Controladores/LoginController"

)

func main() {
	app := iris.New()
	app.HandleDir("/Recursos", "./Recursos")
	app.RegisterView(iris.HTML("./Vistas", ".html").Reload(true))

	app.Get("/", logincontroller.Getlogin)


	app.Run(iris.Addr(":8080"))


}
