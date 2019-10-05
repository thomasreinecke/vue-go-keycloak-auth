# vue-go-keycloak-auth
Test & learning application to utilize a **VueJs** frontend with a **Go** backend speaking with **Keycloak** as identity provider.

**Please note:** The backend is not yet bound to Keycloak, this is to come very shortly. Stay tuned :)

* the Go backend is exposing a REST endpoint /auth (POST) on `localhost:5000`, taking a JSON `{"username":"test@tester.de","password":"test"}` and is delivering a JWT back if the username was `test@tester.de`, else `HTTP/403 Forbidden` 
* the VueJs frontend is providing a login screen on `http://localhost:8080` and is logging the user into the `/main` route on success

![image](https://user-images.githubusercontent.com/35994116/66257232-91477200-e796-11e9-99dc-2f1860dbc539.png)

## Use

* start the backend `cd backend; go run cmd/server/server.go`
* start the frontend `cd frontend; yarn dev` -> go to `http://localhost:8080/` to use the frontend, login with `test@tester.de`

![image](https://user-images.githubusercontent.com/35994116/66257366-7ece3800-e798-11e9-9e5c-f956d6132f94.png)
