package bd

import (
	"context"
	"time"

	"github.com/deinava/testinggolan/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario es la funcion para ver si ya esta registrado*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("golandia")
	col := db.Collection("Usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
