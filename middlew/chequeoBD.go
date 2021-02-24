package middlew

import (
	"net/http"

	"github.com/deinava/testinggolan/bd"
)

/*ChequeoBD middelwere que me permite conocer el estado de la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdidad con la bd", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
