package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/BrangyCastro/twitter-backend/middlew"
	"github.com/BrangyCastro/twitter-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores se seata el puerto, el handler y pongo a escuchar el Servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoDB(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
