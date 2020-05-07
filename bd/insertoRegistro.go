package bd

import (
	"context"
	"time"

	"github.com/BrangyCastro/twitter-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro func que nos permite insertar los datos de los usuarios en la BD
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
