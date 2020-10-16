package admoncontroller

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/kataras/iris/v12"
	"github.com/leekchan/accounting"
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
		espectaculares := admonmodel.ExtraeEspectaculares()
		ctx.ViewData("Espectaculares", espectaculares)
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
		Swal.fire(
			'Eliminado!',
			'Cliente eliminado',
			'success'
		  )
		</script>`)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		Swal.fire(
			'No Eliminado!',
			'Cliente no eliminado',
			'danger'
		  )
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
		Swal.fire('Empleado eliminado');		
		</script>`)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		Swal.fire('Empleado no eliminado');
		</script>`)
	}

	ctx.HTML(htmlcode)
}

//EliminarEspectacular -> Elimina el espectacular de la base de datos
func EliminarEspectacular(ctx iris.Context) {
	id := ctx.PostValue("data")
	var imagenes []string
	espectacular := admonmodel.ObtenerEspectacularPorID(id)
	imagenes = espectacular.Imagenes
	admonmodel.EliminarEspectacularMongo(id)

	for _, v := range imagenes {
		os.Remove(v)
	}
}

//GuardaEspectacular -> Guarda la informacion del espectacular en la base de datos
func GuardaEspectacular(ctx iris.Context) {

	layout := "2006-01-02"
	location, _ := time.LoadLocation("America/Mexico_City")
	var htmlcode string
	var espectacular admonmodel.EspectacularMongo

	espectacular.NumControl = ctx.PostValue("numcontrol")
	espectacular.CostoImpresion, _ = ctx.PostValueFloat64("costoimpreso")
	espectacular.CostoInstalacion, _ = ctx.PostValueFloat64("instalacion")
	espectacular.Calle = ctx.PostValue("calle")
	espectacular.Numero = ctx.PostValue("numero")
	espectacular.Colonia = ctx.PostValue("colonia")
	espectacular.Localidad = ctx.PostValue("localidad")
	espectacular.Municipio = ctx.PostValue("municipio")
	espectacular.Estado = ctx.PostValue("estado")
	espectacular.Latitud = ctx.PostValue("latitud")
	espectacular.Longitud = ctx.PostValue("longitud")
	espectacular.Referencias = ctx.PostValue("referencias")
	espectacular.Ancho, _ = ctx.PostValueFloat64("ancho")
	espectacular.Alto, _ = ctx.PostValueFloat64("alto")
	// material := ctx.PostValue("material")
	// espectacular.Material= ctx.PostValue("material")
	espectacular.Precio, _ = ctx.PostValueFloat64("precio")
	espectacular.Observaciones = ctx.PostValue("observaciones")
	espectacular.Status = ctx.PostValue("status")
	espectacular.Acabados = ctx.PostValue("acabados")
	espectacular.Propietario = ctx.PostValue("nombreprop")
	espectacular.Celular = ctx.PostValue("celular")
	espectacular.Telefono = ctx.PostValue("telefono")

	fechacontratoinicio := ctx.PostValue("iniciocontrato")
	fechacontratofin := ctx.PostValue("fincontrato")

	fmt.Println("Fecha 1", fechacontratoinicio)
	fmt.Println("Fecha 2", fechacontratofin)

	iniciocontrato, _ := time.ParseInLocation(layout, fechacontratoinicio, location)
	fincontrato, _ := time.ParseInLocation(layout, fechacontratofin, location)

	espectacular.ContratoInicio = iniciocontrato
	espectacular.ContratoFin = fincontrato
	espectacular.Monto, _ = ctx.PostValueFloat64("monto")
	espectacular.Folio = ctx.PostValue("folio")
	espectacular.TipoPago = ctx.PostValue("tipopago")
	espectacular.PeriodoPago = ctx.PostValue("periodopago")

	imagen1, header, err := ctx.FormFile("imagen1")
	nombrearchivo1 := header.Filename
	check(err, "Error al seleccionar la imagen 1")

	imagen2, header, err := ctx.FormFile("imagen2")
	nombrearchivo2 := header.Filename
	check(err, "Error al seleccionar la imagen 2")

	imagen3, header, err := ctx.FormFile("imagen3")
	nombrearchivo3 := header.Filename
	check(err, "Error al seleccionar la imagen 3")

	dirPath := "./Recursos/Imagenes/Espectaculares"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Println("el directorio no existe")
		os.MkdirAll(dirPath, 0777)
	} else {
		fmt.Println("el directorio ya existe")
	}
	out, err := os.Create("./Recursos/Imagenes/Espectaculares/" + espectacular.NumControl + "-" + nombrearchivo1)
	check(err, "No se puede crear el archivo revisa los privilegios de escritura.")
	defer out.Close()
	_, err = io.Copy(out, imagen1)
	check(err, "Error al escribir la imagen al directorio 1")
	out2, err2 := os.Create("./Recursos/Imagenes/Espectaculares/" + espectacular.NumControl + "-" + nombrearchivo2)
	_, err2 = io.Copy(out2, imagen2)
	check(err2, "Error al escribir la imagen al directorio 2")
	out3, err3 := os.Create("./Recursos/Imagenes/Espectaculares/" + espectacular.NumControl + "-" + nombrearchivo3)
	_, err3 = io.Copy(out3, imagen3)
	check(err3, "Error al escribir la imagen al directorio 3")

	if nombrearchivo1 == "" {

	} else {
		espectacular.Imagenes = append(espectacular.Imagenes, "Recursos/Imagenes/Espectaculares/"+espectacular.NumControl+"-"+nombrearchivo1)
	}

	if nombrearchivo2 == "" {

	} else {
		espectacular.Imagenes = append(espectacular.Imagenes, "Recursos/Imagenes/Espectaculares/"+espectacular.NumControl+"-"+nombrearchivo2)
	}

	if nombrearchivo3 == "" {

	} else {
		espectacular.Imagenes = append(espectacular.Imagenes, "Recursos/Imagenes/Espectaculares/"+espectacular.NumControl+"-"+nombrearchivo3)
	}

	if admonmodel.GuardarEspectacularMongo(espectacular) {
		htmlcode += fmt.Sprintf(`<script> 
		location.replace("/espectaculares");
		Swal.fire('Espectacular %v agregado');
		</script>`, espectacular.NumControl)
	} else {
		htmlcode += fmt.Sprintf(`<script>
		location.replace("/espectaculares");
		Swal.fire('Espectacular %v no agregado');
	</script>`, espectacular.NumControl)
	}

	ctx.HTML(htmlcode)

}

func check(err error, mensaje string) {
	if err != nil {
		fmt.Println(err)
	}
}

//ObtenerImagenes -> Regresa las imagenes al modal de los espectaculares
func ObtenerImagenes(ctx iris.Context) {

	var htmlcode string

	idespectacular := ctx.PostValue("data")

	espectacular := admonmodel.ObtenerEspectacularPorID(idespectacular)

	htmlcode += fmt.Sprintf(`
	<div id="carouselespectacular" class="carousel slide" data-ride="carousel">
	<ol class="carousel-indicators">`)

	for k, _ := range espectacular.Imagenes {
		if k != 0 {
			htmlcode += fmt.Sprintf(`<li data-target="#carouselespectacular" data-slide-to="%v" class="active"></li>`, k)
		} else {
			htmlcode += fmt.Sprintf(`<li data-target="#carouselespectacular" data-slide-to="%v"</li>`, k)
		}
	}

	htmlcode += fmt.Sprintf(`
	</ol>
	<div class="carousel-inner">`)

	for k, v := range espectacular.Imagenes {

		if k != 0 {
			htmlcode += fmt.Sprintf(`

		<div class="carousel-item">
		<img class="d-block w-100" src="%v" alt="Slide">
	  </div>`, v)

		} else {
			htmlcode += fmt.Sprintf(`

		<div class="carousel-item active">
		<img class="d-block w-100" src="%v" alt="Slide">
	  </div>`, v)
		}

	}

	htmlcode += fmt.Sprintf(`

	</div>
	<a class="carousel-control-prev" href="#carouselespectacular" role="button" data-slide="prev">
	  <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	  <span class="sr-only">Anterior</span>
	</a>
	<a class="carousel-control-next" href="#carouselespectacular" role="button" data-slide="next">
	  <span class="carousel-control-next-icon" aria-hidden="true"></span>
	  <span class="sr-only">Siguiente</span>
	</a>
  </div>`)

	ctx.HTML(htmlcode)

}

//VerificaEspectacular -> Verifica la existencia del numero de control en la base de datos
func VerificaEspectacular(ctx iris.Context) {

	var htmlcode string

	control := ctx.PostValue("data")

	var existe bool

	existe = admonmodel.VerificaEspectacularMongo(control)

	if existe {

		htmlcode += fmt.Sprintf(`
		<script>
		Swal.fire('%v ha sido dado de alta')
		document.getElementById("submitespectacular").disabled = true;
		</script>`, control)

	} else {
		htmlcode += fmt.Sprintf(`
		<script>
		Swal.fire('%v no ha sido dado de alta')
		document.getElementById("submitespectacular").disabled = false;
		</script>`, control)

	}

	ctx.HTML(htmlcode)
}

//GenerarFichaDeCliente -> Genera el archivo PDF por espectacular
func GenerarFichaDeCliente(ctx iris.Context) {

	var htmlcode string
	id := ctx.PostValue("data")

	espectacular := admonmodel.ObtenerEspectacularPorID(id)

	// Construir un pdf.	pdf := gofpdf.New("P", "mm", "A4", "")
	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	var opt gofpdf.ImageOptions
	pdf := gofpdf.New("L", "mm", "Letter", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	// func (f *Fpdf) SetFont(familyStr, styleStr string, size float64)
	// func (f *Fpdf) ImageOptions(imageNameStr string, x, y, w, h float64, flow bool, options ImageOptions, link int, linkStr string)
	// pdf.ImageOptions(`Recursos\Imagenes\nikelogo.jpg`, 0, 100, 20, 0, false, opt, 0, "")
	// pdf.ImageOptions(`Recursos\Imagenes\nikelogo.jpg`, 0, 200, 20, 0, false, opt, 0, "")
	// pdf.ImageOptions(`Recursos\Imagenes\nikelogo.jpg`, 100, 0, 20, 0, false, opt, 0, "")
	// pdf.ImageOptions(`Recursos\Imagenes\pruebamapa.jpg`, 200, 0, 20, 0, false, opt, 0, "")
	// func (f *Fpdf) Cell(w, h float64, txtStr string)
	pdf.ImageOptions(`Recursos\Imagenes\logo.jpg`, 10, 10, 70, 0, false, opt, 0, "")
	pdf.SetDrawColor(168, 170, 173)
	pdf.SetLineWidth(0.5)
	pdf.Line(80, 22, 250, 22)

	pdf.SetFont("Helvetica", "B", 12)
	pdf.ImageOptions(espectacular.Imagenes[0], 30, 35, 70, 0, false, opt, 0, "")
	pdf.SetXY(45, 80)
	pdf.Cell(80, 5, tr("VISTA CERCANA"))
	pdf.ImageOptions(espectacular.Imagenes[1], 105, 35, 70, 0, false, opt, 0, "")
	pdf.SetXY(125, 80)
	pdf.Cell(80, 5, tr("VISTA LEJANA"))
	pdf.ImageOptions(espectacular.Imagenes[2], 180, 35, 70, 0, false, opt, 0, "")
	pdf.SetXY(200, 80)
	pdf.Cell(80, 5, tr("VISTA COMPLETA"))
	pdf.SetFont("Helvetica", "B", 24)
	pdf.SetTextColor(227, 34, 34)
	pdf.SetXY(10, 150)
	pdf.Cell(0, 7, tr(espectacular.NumControl))
	pdf.SetFont("Helvetica", "", 14)
	pdf.SetTextColor(78, 80, 82)
	pdf.Ln(7)
	pdf.Cell(100, 5, tr(espectacular.Calle+" "+espectacular.Numero+", "+espectacular.Colonia+", "+espectacular.Localidad+"."))
	pdf.Ln(5)
	pdf.Cell(50, 5, tr(espectacular.Municipio+", "+espectacular.Estado+"."))
	pdf.Ln(5)
	anchostring := strconv.FormatFloat(espectacular.Ancho, 'f', -1, 64)
	altostring := strconv.FormatFloat(espectacular.Alto, 'f', -1, 64)
	pdf.Cell(50, 5, tr(anchostring+"m x "+altostring+"m"))
	pdf.Ln(5)
	pdf.Cell(50, 5, tr("IMPRESIÓN  : "+ac.FormatMoney(espectacular.CostoImpresion)))
	pdf.Ln(5)
	pdf.Cell(50, 5, tr("INSTALACIÓN: "+ac.FormatMoney(espectacular.CostoInstalacion)))
	pdf.Ln(5)
	pdf.Cell(50, 5, tr(espectacular.Status))
	pdf.SetXY(180, 184)
	pdf.SetTextColor(168, 170, 173)
	pdf.SetFont("Helvetica", "", 9)
	pdf.Cell(80, 4, tr("La Soledad #115, Fracc. Colinas de la Soledad"))
	pdf.SetXY(180, 187)
	pdf.Cell(80, 4, tr("San Felipe del Agua, Oaxaca, Oax. C.P. 68044"))
	pdf.SetXY(180, 190)
	pdf.Cell(80, 4, tr("Tel. (951) 503-82-020, publihome@hotmail.com"))

	pdf.ImageOptions(`Recursos\Imagenes\pruebamapa.jpg`, 150, 95, 100, 0, false, opt, 0, "")
	pdf.SetTextColor(39, 88, 138)
	pdf.SetFont("Helvetica", "B", 16)
	pdf.Text(187, 175, tr(espectacular.Latitud+" , "+espectacular.Longitud))

	fileee := `Recursos\Archivos\Fichas` + espectacular.NumControl + `.pdf`
	err4 := pdf.OutputFileAndClose(fileee)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		htmlcode += fmt.Sprintf(`<script>Descargar('%v');</script>`, espectacular.NumControl)
	}

	ctx.HTML(htmlcode)

}
