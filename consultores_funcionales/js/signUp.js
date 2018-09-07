$(document).ready(function() {
  $('form').submit(function(event) {
      event.preventDefault();
      var attribute = {
        Name : 'email',
        Value : $('#email').val()
      };
      var attributeEmail = new AmazonCognitoIdentity.CognitoUserAttribute(attribute);
      var attributeList = [];
      attributeList.push(attributeEmail);
      var username = $('#email').val();
      var password = $('#password').val();
      userPool.signUp(username, password, attributeList, null, function(err, result) {
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
