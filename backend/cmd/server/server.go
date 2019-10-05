package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/thomasreinecke/vue-go-keycloak-auth/backend/internal/middleware"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	w.Write([]byte("foo"))
}

func main() {

	fmt.Println("Starting finstack/auth REST service...")
	// initialize router
	router := mux.NewRouter()

	// configure the router
	router.Use(loggingMiddleware)

	//endpoints
	router.HandleFunc("/auth", middleware.Authenticate).Methods("POST")

	// added CORS for heroku deploy
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":5000", handler))
}
