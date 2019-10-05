package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v3"
	"github.com/thomasreinecke/vue-go-keycloak-auth/backend/internal/user"
)

// Authenticate tries to authenticate a user by its given username and password
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.LoginUser
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("decoded to user", u)

	// client := gocloak.NewClient("http://localhost:38080")
	// token, err := client.LoginAdmin("admin", "admin", "master")
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Something wrong with the credentials or url")
	// }
	// fmt.Println("Token: ", token)

	client := gocloak.NewClient("http://localhost:38080")
	// token, err := client.Login("vue-go-keycloak-auth", "", "master", u.Username, u.Password)
	token, err := client.Login("admin-cli", "", "master", u.Username, u.Password)

	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		fmt.Println("token: ", token)
		json.NewEncoder(w).Encode(token)
	}

	//  TRADITIONAL WAY WITHOUT KEYCLOAK

	// if u.Username == "test@tester.de" {
	// 	fmt.Println("Great stuff, authenticated !")

	// 	// Create the Claims
	// 	jwtUser := user.JwtUserToken{
	// 		user.WebUser{
	// 			"0",
	// 			"test@tester.de",
	// 			"tester",
	// 			true,
	// 		},
	// 		jwt.StandardClaims{
	// 			ExpiresAt: 15000,
	// 			Issuer:    "test",
	// 		},
	// 	}

	// 	mySigningKey := []byte("AllYourBase")
	// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtUser)
	// 	ss, err := token.SignedString(mySigningKey)
	// 	if err == nil {
	// 		response := user.JwtUserTokenResponse{
	// 			ss,
	// 		}
	// 		json.NewEncoder(w).Encode(response)
	// 	}
	// } else {
	// 	http.Error(w, "Forbidden", http.StatusForbidden)
	// }
}
