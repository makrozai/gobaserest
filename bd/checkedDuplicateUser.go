package bd

import (
	"context"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckedDuplicateUser recibe un email de parametro y chequea si esta en la BD
func CheckedDuplicateUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var user models.User

	err := col.FindOne(ctx, condition).Decode(&user)
	ID := user.ID.Hex()

	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}
