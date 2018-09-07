  var CognitoUserPool = AmazonCognitoIdentity.CognitoUserPool;

  var poolData = {
    UserPoolId : 'us-east-1_2jFMqHgQB',
    ClientId : '4r058570sc99gabq0hi1tdo4f4'
  }

  var userPool = new CognitoUserPool(poolData);
