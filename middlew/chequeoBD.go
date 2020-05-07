package middlew

import (
	"net/http"

	"github.com/BrangyCastro/twitter-backend/bd"
)

// ChequeoDB middlew que verifica si la conexion de la BD esta conectada
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
