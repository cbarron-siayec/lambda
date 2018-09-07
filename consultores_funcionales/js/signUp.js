$(document).ready(function() {
  $('form').submit(function(event) {
      event.preventDefault();
      var username = $('#email').val();
      var password = $('#password').val();
      userPool.signUp(username, password, null, null, function(err, result) {
        if (err) {
          $('#error').html("Error: "+err);
          console.log(err);
          return;
        }
        console.log(result);
        window.location.href = "./validar-correo-electronico.html";
      });
  });
});
