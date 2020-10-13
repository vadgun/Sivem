// let inicio = document.getElementById("indexit");
// inicio.classList.add("selected");

let clientes = document.getElementById("clientesit");
let empleados = document.getElementById("empleadosit");
let espectaculares = document.getElementById("espectacularesit");

$('.dropify').dropify();
function BuscaClientes() {
    inputvalue = document.getElementById("buscadorcliente").value;
    console.log(inputvalue+" AQUI");
}

var input = document.getElementById("buscadorcliente");
console.log(input);

if (input == null) {

}else{ 
  input.addEventListener("keyup", function(event) {
    if (event.key === "Enter") {
      event.preventDefault();
      BuscaClientes();
    }
  });
}

function EditarCliente(data){

  $.ajax({
    url: '/obtenercliente',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#formeditarclientes").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Datos de edicion de cliente obtenidos")
        
    }
});

$("#ModalEditarClientes").modal("show");

}

function EliminarCliente(data){

var r = confirm("Elimnar cliente, ¿Desea continuar?");
if (r == true) {
  $.ajax({
    url: '/eliminarcliente',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#formeditarclientes").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Datos de edicion de cliente obtenidos")
        
    }
});  
} else {
  return false;
}

}

function EditarEmpleado(data){

  $.ajax({
    url: '/obtenerempleado',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#formeditarempleado").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Datos de edicion de empleado obtenidos")
        
    }
});

$("#ModalEditarEmpleados").modal("show");

}

function EliminarEmpleado(data){

  var r = confirm("Elimnar empleado, ¿Desea continuar?");
if (r == true) {
  $.ajax({
    url: '/eliminarempleado',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#formeditarempleado").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Datos de edicion de cliente obtenidos")
        
    }
});  
} else {
  return false;
}
}

function Ubicaciones(data){
  if (data == ""){
    return false
  }else{
    alert(data);
  }

}

function EditarEspectacular(data){

}
function EliminarEspectacular(data){

}
function VerImagenesEspectacular(data){
  
  $.ajax({
    url: '/imagenesespectacular',
    data: { data: data },
    type: 'POST',
    dataType: 'html',
    success: function(result) {
        console.log("Operacion Realizada con Exito");
        $("#ContainerImagenesEspectacular").html(result);
    },
    error: function(xhr, status) {
        console.log("Error en la consulta")
    },
    complete: function(xhr, status) {
        console.log("Datos de imagenes para modal")
        
    }
});  

$("#ModalImagenesEspectacular").modal("show");


}