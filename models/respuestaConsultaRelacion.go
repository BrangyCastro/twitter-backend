package models

// RespuestaConsultaRelacion TIENE EL TRUE O FALSE QYE SE OBTIENE DE CONSULTAR LA RELACION ENTRE 2 USUARIOS
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}