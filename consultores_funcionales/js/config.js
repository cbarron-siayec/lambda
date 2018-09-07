  var CognitoUserPool = AmazonCognitoIdentity.CognitoUserPool;

  var poolData = {
    UserPoolId : 'us-east-1_D6OhQx0nt',
    ClientId : '1hd69fdkjh5on8t252mjfl26qg'
  }

  var userPool = new CognitoUserPool(poolData);
