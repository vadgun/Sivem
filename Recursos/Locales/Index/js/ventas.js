const facturaSelect = document.querySelector("#factura")
facturaSelect.addEventListener("change", function(e){
         if(this.value == "si"){
                 window.iva.classList.add("d-block")
         }else{
                 window.iva.classList.remove("d-block")

         }
})