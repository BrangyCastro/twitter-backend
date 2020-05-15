package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BrangyCastro/twitter-backend/bd"
)

// LeoTweetsRelacion lee los tweets de todos nbuestros seguidores
func LeoTweetsRelacion(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetSeguidores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error a leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
