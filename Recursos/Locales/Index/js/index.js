// let inicio = document.getElementById("indexit");
// inicio.classList.add("selected");

// let clientes = document.getElementById("clientesit");
// clientes.classList.add("selected");

function BuscaClientes() {
    inputvalue = document.getElementById("buscadorcliente").value;
    console.log(inputvalue);
}

var input = document.getElementById("buscadorcliente");
input.addEventListener("keyup", function(event) {
    if (event.key === "Enter") {
      event.preventDefault();
      BuscaClientes();
    }
  });