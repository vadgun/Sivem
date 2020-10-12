const button = document.getElementById('menu');
const nav =  document.querySelector('.sidenav');

$('#menu').click(function(e){
    e.preventDefault();
    $('.sidenav').toggle('slide')
})