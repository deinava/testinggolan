package routers

import (
	"encoding/json"
	"net/http"

	"github.com/deinava/testinggolan/bd"
	"github.com/deinava/testinggolan/models"
)

/*Registro es la fncion es para crear el usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El passowrd es menor a 6", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Email ya registrado", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al guardar la informacion del usuario"+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usurio", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
