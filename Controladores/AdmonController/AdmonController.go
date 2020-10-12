package admoncontroller

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
	sessioncontroller "github.com/vadgun/Sivem/Controladores/SessionController"
	admonmodel "github.com/vadgun/Sivem/Modelos/AdmonModel"
	indexmodel "github.com/vadgun/Sivem/Modelos/IndexModel"
)

//Clientes -> Regresa la pagina de inicio
func Clientes(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		clientes := admonmodel.ExtraeClientes()
		ctx.ViewData("Clientes", clientes)
		if err := ctx.View("Clientes.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Empleados -> Regresa la pagina de inicio
func Empleados(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		empleados := admonmodel.ExtraeEmpleados()
		ctx.ViewData("Empleados", empleados)
		if err := ctx.View("Empleados.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Empresas -> Regresa la pagina de inicio
func Empresas(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Empresas.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Medios -> Regresa la pagina de inicio
func Medios(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Medios.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Materiales -> Regresa la pagina de inicio
func Materiales(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Materiales.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Espectaculares -> Regresa la pagina de inicio
func Espectaculares(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Espectaculares.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//EspectacularesNuevo -> Regresa la vista autorizada para el formulario de captura de un espectacular y las consultas que este implica
func EspectacularesNuevo(ctx iris.Context) {

	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("EspectacularesNuevo.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}

}

//Ventas -> Regresa la pagina de inicio
func Ventas(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Ventas.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//Vallas -> Regresa la pagina de inicio
func Vallas(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("Vallas.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//VallasMoviles -> Regresa la pagina de inicio
func VallasMoviles(ctx iris.Context) {
	var usuario indexmodel.MongoUser
	var autorizado bool
	autorizado2, _ := sessioncontroller.Sess.Start(ctx).GetBoolean("Autorizado")

	if autorizado2 == false {
		usuario.Key = ctx.PostValue("pass")
		usuario.Usuario = ctx.PostValue("usuario")
		autorizado, usuario = indexmodel.VerificarUsuario(usuario)
		if autorizado {
			sessioncontroller.Sess.Start(ctx).Set("Autorizado", true)
			sessioncontroller.Sess.Start(ctx).Set("UserID", usuario.ID.Hex())
		}
	}

	if autorizado || autorizado2 {
		userOn := indexmodel.GetUserOn(sessioncontroller.Sess.Start(ctx).GetString("UserID"))
		ctx.ViewData("Usuario", userOn)
		if err := ctx.View("VallasMoviles.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	} else {
		ctx.Redirect("/login", iris.StatusSeeOther)
	}
}

//F O R M U L A R I O    D E   C L I E N T E

//GuardaClientes -> Guarda la captura de los datos del cliente
func GuardaClientes(ctx iris.Context) {

	var cliente admonmodel.ClienteMongo
	var htmlcode string

	cliente.Nombre = strings.ToUpper(ctx.PostValue("nombrecliente"))
	cliente.ApellidoP = strings.ToUpper(ctx.PostValue("apaterno"))
	cliente.ApellidoM = strings.ToUpper(ctx.PostValue("amaterno"))
	cliente.Sexo = strings.ToUpper(ctx.PostValue("sexo"))
	cliente.Empresa = strings.ToUpper(ctx.PostValue("empresa"))
	cliente.Puesto = strings.ToUpper(ctx.PostValue("puesto"))
	cliente.CorreoEMP = strings.ToUpper(ctx.PostValue("correoemp"))
	cliente.Correo = strings.ToUpper(ctx.PostValue("correo"))
	cliente.Calle = strings.ToUpper(ctx.PostValue("calle"))
	cliente.Numero = strings.ToUpper(ctx.PostValue("numero"))
	cliente.Colonia = strings.ToUpper(ctx.PostValue("colonia"))
	cliente.CodigoP = strings.ToUpper(ctx.PostValue("codpostal"))
	cliente.Municipio = strings.ToUpper(ctx.PostValue("municipio"))
	cliente.Estado = strings.ToUpper(ctx.PostValue("estado"))
	cliente.Ciudad = strings.ToUpper(ctx.PostValue("ciudad"))
	cliente.Telefono = strings.ToUpper(ctx.PostValue("telefono"))
	cliente.TelefonoA = strings.ToUpper(ctx.PostValue("telefonob"))
	cliente.SitioWeb = strings.ToUpper(ctx.PostValue("sitioweb"))
	cliente.Observaciones = strings.ToUpper(ctx.PostValue("observaciones"))

	if admonmodel.GuardarClienteMongo(cliente) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente Guardado %v");
		location.replace("/clientes");
		</script>`, cliente.Nombre)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente no guardado ");
		location.replace("/clientes");
		</script>`)
	}

	ctx.HTML(htmlcode)

}

//EditandoClientes -> Guarda el Cliente en la base de datos
func EditandoClientes(ctx iris.Context) {
	var cliente admonmodel.ClienteMongo
	var htmlcode string
	idcliente := ctx.PostValue("id2")
	cliente.Nombre = strings.ToUpper(ctx.PostValue("nombrecliente2"))
	cliente.ApellidoP = strings.ToUpper(ctx.PostValue("apaterno2"))
	cliente.ApellidoM = strings.ToUpper(ctx.PostValue("amaterno2"))
	cliente.Sexo = strings.ToUpper(ctx.PostValue("sexo2"))
	cliente.Empresa = strings.ToUpper(ctx.PostValue("empresa2"))
	cliente.Puesto = strings.ToUpper(ctx.PostValue("puesto2"))
	cliente.CorreoEMP = strings.ToUpper(ctx.PostValue("correoemp2"))
	cliente.Correo = strings.ToUpper(ctx.PostValue("correo2"))
	cliente.Calle = strings.ToUpper(ctx.PostValue("calle2"))
	cliente.Numero = strings.ToUpper(ctx.PostValue("numero2"))
	cliente.Colonia = strings.ToUpper(ctx.PostValue("colonia2"))
	cliente.CodigoP = strings.ToUpper(ctx.PostValue("codpostal2"))
	cliente.Municipio = strings.ToUpper(ctx.PostValue("municipio2"))
	cliente.Estado = strings.ToUpper(ctx.PostValue("estado2"))
	cliente.Ciudad = strings.ToUpper(ctx.PostValue("ciudad2"))
	cliente.Telefono = strings.ToUpper(ctx.PostValue("telefono2"))
	cliente.TelefonoA = strings.ToUpper(ctx.PostValue("telefonob2"))
	cliente.SitioWeb = strings.ToUpper(ctx.PostValue("sitioweb2"))
	cliente.Observaciones = strings.ToUpper(ctx.PostValue("observaciones2"))

	if admonmodel.EditarClienteMongo(idcliente, cliente) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente %v Editado");
		location.replace("/clientes");
		</script>`, cliente.Nombre)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente no editado ");
		location.replace("/clientes");
		</script>`)
	}

	ctx.HTML(htmlcode)

}

//ObtenerCliente -> Regresa el Cliente al Modal de la Peticion de Editar
func ObtenerCliente(ctx iris.Context) {

	var htmlcode string

	clienteid := ctx.PostValue("data")

	cliente := admonmodel.ExtraeClientePorId(clienteid)

	htmlcode += fmt.Sprintf(`
	<div class="row">
                        <div class="col">
							<input type="hidden" value="%v" name="id2" id="id2">		
                            <input type="text" class="form-control textomayus" placeholder="Nombre" required value="%v" id="nombrecliente2" name="nombrecliente2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Apellido Paterno" required value="%v" id="apaterno2" name="apaterno2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Apellido Materno" value="%v" id="amaterno2" name="amaterno2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Sexo" value="%v" id="sexo2" name="sexo2">
                        </div>
                   </div>
				   <br>`, cliente.ID.Hex(), cliente.Nombre, cliente.ApellidoP, cliente.ApellidoM, cliente.Sexo)
	htmlcode += fmt.Sprintf(`
	
                   <div class="row">
                    <div class="col">
                        <input type="text" class="form-control textomayus" placeholder="Empresa" value="%v" id="empresa2" name="empresa2">
                    </div>
                    <div class="col">
                        <input type="text" class="form-control textomayus" placeholder="Puesto" value="%v" id="puesto2" name="puesto2">
                    </div>
                    <div class="col">
                        <input type="text" class="form-control textomayus" placeholder="Correo Empresarial" value="%v" id="correoemp2" name="correoemp2">
                    </div>
                    </div>
					<br>`, cliente.Empresa, cliente.Puesto, cliente.CorreoEMP)
	htmlcode += fmt.Sprintf(`
	
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Correo Electronico" required value="%v" id="correo2" name="correo2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Calle" required value="%v" id="calle2" name="calle2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Numero" required value="%v" id="numero2" name="numero2">
                        </div>

                    </div>
                    <br>`, cliente.Correo, cliente.Calle, cliente.Numero)
	htmlcode += fmt.Sprintf(`
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Colonia" required value="%v" id="colonia2" name="colonia2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Codigo Postal" value="%v" id="codpostal2" name="codpostal2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Municipio" required value="%v" id="municipio2" name="municipio2">
                        </div>
                    </div>
					<br>`, cliente.Colonia, cliente.CodigoP, cliente.Municipio)
	htmlcode += fmt.Sprintf(`
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Estado" required value="%v" id="estado2" name="estado2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Ciudad" required value="%v" id="ciudad2" name="ciudad2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Teléfono" value="%v" id="telefono2" name="telefono2">
                        </div>
                    </div>
					<br>`, cliente.Estado, cliente.Ciudad, cliente.Telefono)
	htmlcode += fmt.Sprintf(`					
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Teléfono Alternativo" value="%v" id="telefonob2" name="telefonob2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Sitio Web" value="%v" id="sitioweb2" name="sitioweb2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Observaciones" value="%v" id="observaciones2" name="observaciones2">
                        </div>
                    </div>
					<br>`, cliente.TelefonoA, cliente.SitioWeb, cliente.Observaciones)
	htmlcode += fmt.Sprintf(`	
                    <div class="row">
                        <div class="col-md-10">
                        </div>
                        <div class="col-sm-1">    
                            <button type="submit" class="btn btn-dark" >Guardar</button>
                        </div>
                    </div>`)

	ctx.HTML(htmlcode)

}

//EliminarCliente -> Elimina el cliente de la base de datos
func EliminarCliente(ctx iris.Context) {

	id := ctx.PostValue("data")
	var htmlcode string

	if admonmodel.EliminarClienteMongo(id) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente Eliminado");
		location.replace("/clientes");
		</script>`)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Cliente no eliminado ");
		location.replace("/clientes");
		</script>`)
	}

	ctx.HTML(htmlcode)

}

//A L T A   D E    E M P L E A D O S
//GuardaEmpleados -> Guarda Empleado en la Base de datos
func GuardaEmpleados(ctx iris.Context) {
	var htmlcode string

	var empleado admonmodel.EmpleadoMongo

	empleado.Nombre = strings.ToUpper(ctx.PostValue("nombreempleado"))
	empleado.ApellidoP = strings.ToUpper(ctx.PostValue("apaterno"))
	empleado.ApellidoM = strings.ToUpper(ctx.PostValue("amaterno"))
	empleado.Sexo = strings.ToUpper(ctx.PostValue("sexo"))
	empleado.Licencia = strings.ToUpper(ctx.PostValue("licencia"))
	empleado.Cargo = strings.ToUpper(ctx.PostValue("cargo"))
	empleado.Correo = strings.ToUpper(ctx.PostValue("correo"))
	empleado.Calle = strings.ToUpper(ctx.PostValue("calle"))
	empleado.Numero = strings.ToUpper(ctx.PostValue("numero"))
	empleado.Colonia = strings.ToUpper(ctx.PostValue("colonia"))
	empleado.CodigoP = strings.ToUpper(ctx.PostValue("codpostal"))
	empleado.Municipio = strings.ToUpper(ctx.PostValue("municipio"))
	empleado.Estado = strings.ToUpper(ctx.PostValue("estado"))
	empleado.Ciudad = strings.ToUpper(ctx.PostValue("ciudad"))
	empleado.Telefono = strings.ToUpper(ctx.PostValue("telefono"))
	empleado.Referencias = strings.ToUpper(ctx.PostValue("referencias"))
	empleado.Control = strings.ToUpper(ctx.PostValue("control"))

	if admonmodel.GuardaEmpleadoMongo(empleado) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado %v Guardado");
		location.replace("/empleados");
		</script>`, empleado.Nombre)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado no guardado");
		location.replace("/empleados");
		</script>`)
	}

	ctx.HTML(htmlcode)

}

//EditandoEmpleados -> Edita los empleados en la BAse de datos
func EditandoEmpleados(ctx iris.Context) {

	var htmlcode string

	var empleado admonmodel.EmpleadoMongo

	idempleado := ctx.PostValue("id2")
	empleado.Nombre = strings.ToUpper(ctx.PostValue("nombreempleado2"))
	empleado.ApellidoP = strings.ToUpper(ctx.PostValue("apaterno2"))
	empleado.ApellidoM = strings.ToUpper(ctx.PostValue("amaterno2"))
	empleado.Sexo = strings.ToUpper(ctx.PostValue("sexo2"))
	empleado.Licencia = strings.ToUpper(ctx.PostValue("licencia2"))
	empleado.Cargo = strings.ToUpper(ctx.PostValue("cargo2"))
	empleado.Correo = strings.ToUpper(ctx.PostValue("correo2"))
	empleado.Calle = strings.ToUpper(ctx.PostValue("calle2"))
	empleado.Numero = strings.ToUpper(ctx.PostValue("numero2"))
	empleado.Colonia = strings.ToUpper(ctx.PostValue("colonia2"))
	empleado.CodigoP = strings.ToUpper(ctx.PostValue("codpostal2"))
	empleado.Municipio = strings.ToUpper(ctx.PostValue("municipio2"))
	empleado.Estado = strings.ToUpper(ctx.PostValue("estado2"))
	empleado.Ciudad = strings.ToUpper(ctx.PostValue("ciudad2"))
	empleado.Telefono = strings.ToUpper(ctx.PostValue("telefono2"))
	empleado.Referencias = strings.ToUpper(ctx.PostValue("referencias2"))
	empleado.Control = strings.ToUpper(ctx.PostValue("control2"))

	if admonmodel.EditarEmpleadoMongo(idempleado, empleado) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado %v Editado");
		location.replace("/empleados");
		</script>`, empleado.Nombre)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado no editado ");
		location.replace("/empleados");
		</script>`)
	}

	ctx.HTML(htmlcode)

}

//ObtenerEmpleado -> Regresa el Empleado al Modal de la Peticion de Editar
func ObtenerEmpleado(ctx iris.Context) {

	var htmlcode string

	empleadoid := ctx.PostValue("data")

	empleado := admonmodel.ExtraeEmpleadoPorId(empleadoid)

	htmlcode += fmt.Sprintf(`
	<div class="row">
	<div class="col">
	<input type="text" class="form-control textomayus" placeholder="# Control" value="%v" id="control2" name="control2">
</div>
                        <div class="col">
							<input type="hidden" value="%v" name="id2" id="id2">		
                            <input type="text" class="form-control textomayus" placeholder="Nombre" required value="%v" id="nombreempleado2" name="nombreempleado2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Apellido Paterno" required value="%v" id="apaterno2" name="apaterno2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Apellido Materno" value="%v" id="amaterno2" name="amaterno2">
                        </div>

                   </div>
				   <br>`, empleado.Control, empleado.ID.Hex(), empleado.Nombre, empleado.ApellidoP, empleado.ApellidoM)
	htmlcode += fmt.Sprintf(`
	
                   <div class="row">
				   <div class="col">
				   <input type="text" class="form-control textomayus" placeholder="Sexo" value="%v" id="sexo2" name="sexo2">
			   </div>
                    <div class="col">
                        <input type="text" class="form-control textomayus" placeholder="Puesto" value="%v" id="cargo2" name="cargo2">
                    </div>
                    <div class="col">
                        <input type="text" class="form-control textomayus" placeholder="Licencia" value="%v" id="licencia2" name="licencia2">
                    </div>
                    </div>
					<br>`, empleado.Sexo, empleado.Cargo, empleado.Licencia)
	htmlcode += fmt.Sprintf(`
	
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Correo Electronico" required value="%v" id="correo2" name="correo2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Calle" required value="%v" id="calle2" name="calle2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Numero" required value="%v" id="numero2" name="numero2">
                        </div>

                    </div>
                    <br>`, empleado.Correo, empleado.Calle, empleado.Numero)
	htmlcode += fmt.Sprintf(`
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Colonia" required value="%v" id="colonia2" name="colonia2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Codigo Postal" value="%v" id="codpostal2" name="codpostal2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Municipio" required value="%v" id="municipio2" name="municipio2">
                        </div>
                    </div>
					<br>`, empleado.Colonia, empleado.CodigoP, empleado.Municipio)
	htmlcode += fmt.Sprintf(`
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Estado" required value="%v" id="estado2" name="estado2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Ciudad" required value="%v" id="ciudad2" name="ciudad2">
                        </div>
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Teléfono" value="%v" id="telefono2" name="telefono2">
                        </div>
                    </div>
					<br>`, empleado.Estado, empleado.Ciudad, empleado.Telefono)
	htmlcode += fmt.Sprintf(`					
                    <div class="row">
                        <div class="col">
                            <input type="text" class="form-control textomayus" placeholder="Referencias" value="%v" id="referencias2" name="referencias2">
                        </div>
                    </div>
					<br>`, empleado.Referencias)
	htmlcode += fmt.Sprintf(`	
                    <div class="row">
                        <div class="col-md-10">
                        </div>
                        <div class="col-sm-1">    
                            <button type="submit" class="btn btn-dark" >Guardar</button>
                        </div>
                    </div>`)

	ctx.HTML(htmlcode)

}

//EliminarEmpleado -> Elimina un empleado de la base de datos
func EliminarEmpleado(ctx iris.Context) {
	id := ctx.PostValue("data")
	var htmlcode string

	if admonmodel.EliminarEmpleadoMongo(id) {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado Eliminado");
		location.replace("/empleados");
		</script>`)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		alert("Empleado no eliminado ");
		location.replace("/empleados");
		</script>`)
	}

	ctx.HTML(htmlcode)
}
