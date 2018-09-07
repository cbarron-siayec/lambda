$(document).ready(function() {
  $('#validateCode').submit(function(event) {
      event.preventDefault();
      var codigo = $('#codigo').val();
      userPool.signUp(username, password, null, null, function(err, result) {
        if (err) {
          $('#error').html("Error: "+err);
          console.log(err);
          return;
        }
        console.log(result);
        window.location.href = "../index.html";
      });
  });
});
