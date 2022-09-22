package bd

import (
	"context"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister es la parada final con la database para insertar los datos del usuario
func InsertRegister(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("users")

	user.Password, _ = EncriptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
