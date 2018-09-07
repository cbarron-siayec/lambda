$(document).ready(function() {
  $('#validateCode').submit(function(event) {
      event.preventDefault();
      var codigo = $('#codigo').val();
      var userData = {
        Username : $('#email').val(),
        Pool : userPool
    };
    var cognitoUser = new AmazonCognitoIdentity.CognitoUser(userData);
    cognitoUser.confirmRegistration(codigo, true, function(err, result) {
        if (err) {
            $('#error').html("Error: "+err);
            return;
        }
        window.location.href = "../index.html";
    });
  });
});
function checkLogin(){
  var code = 0;
  var mail = $('#email').val();
  var code = $('#codigo').val();
  var domain = mail.split("@")[1];
  if(mail !="" && code != 0 && domain =="grupo-siayec.com.mx"){
    $('#validateSubmit').prop('disabled', false);
  }
  else {
    $('#validateSubmit').prop('disabled', true);
  }
}
