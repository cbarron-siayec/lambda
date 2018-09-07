//var cognitoUser = new AWSCognito.CognitoIdentityServiceProvider.CognitoUser({
//          Username : username,
//          Pool : userPool
//});
  var cognitoUser = userPool.getCurrentUser();
  if(cognitoUser){
  }
  else {
    window.location.href = "../index.html";
    logout();
  }
