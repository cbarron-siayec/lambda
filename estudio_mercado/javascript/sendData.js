function test(){
  var API_URL="https://7stqkd972h.execute-api.us-west-1.amazonaws.com/DEV/nuevoregistro";

  $.ajax({
      type: 'POST',
      url: API_URL,
      data: `{"nombreColaborador":"Angela Merkel","fecha":"2018/08/30"}`,
      contentType: "application/json",

      success: function(data){
        location.reload();
      }
  });
  return false;
}
