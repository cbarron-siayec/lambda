$(document).ready(function() {
  $('form').submit(function(event) {
      event.preventDefault();
      var username = $('#email').val();
      var password = $('#password').val();
      var authenticationDetails = new AWSCognito.CognitoIdentityServiceProvider.AuthenticationDetails({
               Username : username,
               Password : password
      });

      var cognitoUser = new AWSCognito.CognitoIdentityServiceProvider.CognitoUser({
                Username : username,
                Pool : userPool
      });

      cognitoUser.authenticateUser(authenticationDetails, {
        onSuccess: function (result) {
          window.location.href = "./html/recuperar-contraseña.html";
          console.log(result);
        },
        onFailure: function(err) {
          $('#error').html("Error: "+err);
          console.log(err);
        }
      });
  });
});
