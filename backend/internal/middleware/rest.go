package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/thomasreinecke/vue-go-keycloak-auth/backend/internal/user"
)

// Authenticate tries to authenticate a user by its given username and password
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.LoginUser
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("decoded to user", u)

	if u.Username == "test@tester.de" {
		fmt.Println("Great stuff, authenticated !")

		// Create the Claims
		jwtUser := user.JwtUserToken{
			user.WebUser{
				"0",
				"test@tester.de",
				"tester",
				true,
			},
			jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "test",
			},
		}

		mySigningKey := []byte("AllYourBase")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtUser)
		ss, err := token.SignedString(mySigningKey)
		if err == nil {
			response := user.JwtUserTokenResponse{
				ss,
			}
			json.NewEncoder(w).Encode(response)
		}
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
