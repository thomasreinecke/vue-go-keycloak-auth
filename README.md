# vue-go-keycloak-auth
Test & learning application to utilize a **VueJs** frontend with a **Go** backend speaking with **Keycloak** as identity provider.

* the Go backend is exposing a REST endpoint /auth (POST) on `localhost:5000`, taking a JSON `{"username":"test@tester.de","password":"test"}` and is delivering a JWT back if the username was `test@tester.de`, else `HTTP/403 Forbidden` 
* the VueJs frontend is providing a login screen on `http://localhost:8080` and is logging the user into the `/main` route on success

![image](https://user-images.githubusercontent.com/35994116/66257232-91477200-e796-11e9-99dc-2f1860dbc539.png)

## Installation & How to Use

* pull the `jboss/keycloak` image `docker pull jboss/keycloak`
* start the image `docker run -p 38080:8080 --name keycloak -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin jboss/keycloak`
* login to the keycloak admin console on `http://localhost:38080/auth/` > Users > Add User > create the `test@tester.de` user
![image](https://user-images.githubusercontent.com/35994116/66257908-19317a00-e79f-11e9-9dea-dfbb4528f4c2.png)
* when the user is successfully created, on the **Credentials** tab, set `test` as new password for this user (make sure you switch **Temporary** off)
![image](https://user-images.githubusercontent.com/35994116/66257589-61e73400-e79b-11e9-92ac-8a61c88b2a48.png)

* start the backend `cd backend; go run cmd/server/server.go`
* prep frontend node modules `cd frontend; yarn`
* start the frontend `yarn dev` (in frontend) -> go to `http://localhost:8080/` to use the frontend, login with `test@tester.de`

![image](https://user-images.githubusercontent.com/35994116/66257366-7ece3800-e798-11e9-9e5c-f956d6132f94.png)
