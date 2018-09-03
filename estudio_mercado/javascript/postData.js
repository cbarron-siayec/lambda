$(document).ready(function() {

  $('form').submit(function(event) {

    var API_URL="https://7stqkd972h.execute-api.us-west-1.amazonaws.com/DEV/nuevoregistro";
    event.preventDefault();
    $.ajax({
        type: 'POST',
        url: API_URL,
        data: JSON.stringify({
          "nombreColaborador":$('#nombreColaborador').val(),
          "fecha":$('#fecha').val(),
          "presupuesto":$('input:radio[name=presupuesto]:checked').val(),
          "horasDia":$('input:radio[name=horasDia]:checked').val(),
          "personalExclusivo":$('input:radio[name=personalExclusivo]:checked').val(),
          "manejoDatos":$('input:radio[name=manejoDatos]:checked').val(),
          "duplicarTrabajo":$('input:radio[name=duplicarTrabajo]:checked').val(),
          "infraestructuraNube":$('input:radio[name=infraestructuraNube]:checked').val(),
          "colaboradoresAcceso":$('input:radio[name=colaboradoresAcceso]:checked').val(),
          "utilidad":$('#utilidad').val(),
          "seguridadPublica":$('#seguridadPublica').val(),
          "callesParquesJardines":$('#callesParquesJardines').val(),
          "aguaDrenaje":$('#aguaDrenaje').val(),
          "serviciosLimpia":$('#serviciosLimpia').val(),
          "alumbrado":$('#alumbrado').val(),
          "tramites":$('#tramites').val(),
          "mercados":$('#mercados').val(),
          "administracion":$('#administracion').val(),
          "comentarios":$('#comentarios').val(),
        }),
        success: function(data){
          location.reload();
        }
    });
    return false;

  });

});
