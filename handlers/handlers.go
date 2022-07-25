package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/JosseMontano/InstagramReactGolang/middlew"
	"github.com/JosseMontano/InstagramReactGolang/routers"
)

//drivers this function allow set the port, the handler y listen the server
func Drivers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verPerfil", middlew.CheckBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("POST")



	
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
