# cookie-session
the cooke&session practise

Make a sign-in request with the appropriate credentials
POST http://localhost:8080/signin
{"username":"user2","password":"password2"}
GET http://localhost:8080/signin?Username=user2&Password=password2

And now try to get welcome message
GET http://localhost:8080/welcome

Refresh rout to get a new session_token
GET http://localhost:8080/refresh

Finally, logout to clear session data
GET http://localhost:8080/logout
