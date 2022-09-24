# cookie-session
The cooke&session example

* Make a sign-in request with the appropriate credentials  
`POST http://localhost:8080/signin`{"username":"user2","password":"password2"}  
`GET http://localhost:8080/signin?Username=user2&Password=password2`  
* Get welcome message
`GET http://localhost:8080/welcome`

* Refresh rout to get a new session_token
`GET http://localhost:8080/refresh`

* Logout to clear session data
`GET http://localhost:8080/logout`
