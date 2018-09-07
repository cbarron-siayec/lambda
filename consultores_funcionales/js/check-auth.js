  var cognitoUser = userPool.getCurrentUser();
  var windowCheck;
  cognitoUser.getSession(function sessionCallback(err, session) {
                  if (err) {
                      windowCheck = "";
                  } else if (!session.isValid()) {
                      windowCheck = "";
                  } else {
                      windowCheck = session.getIdToken().getJwtToken();
                  }
  });
if (windowCheck == "") {
  window.location.href = "../index.html";
  logout();
}
