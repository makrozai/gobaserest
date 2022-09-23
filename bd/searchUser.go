package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchUser(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("users")

	var user models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&user)
	user.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado" + err.Error())
		return user, err
	}
	return user, nil
}
