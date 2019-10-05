
# backend of vue-go-keycloak-auth 
this project provides a RESTful service exposing a keycloak user authentication


## cmd/server/server
REST service bound to `::5000` thats providing the REST service at `http://localhost:5000/sushi`

## Installation & How to Use
```
cd backend; go run cmd/server/server.go
```

the **auth** service exposes the following REST interface:
* POST `/auth` authenticates a user (structure is `{"username":"test@tester.de","password":"test"}`
