package bd

import (
	"context"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateUser permite modificar el perfil del usuario
func UpdateUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("users")

	userRequest := make(map[string]interface{})
	if len(user.Name) > 0 {
		userRequest["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		userRequest["lastName"] = user.LastName
	}
	userRequest["dateBirth"] = user.DateBirth
	if len(user.Avatar) > 0 {
		userRequest["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		userRequest["banner"] = user.Banner
	}
	if len(user.Description) > 0 {
		userRequest["description"] = user.Description
	}
	if len(user.Location) > 0 {
		userRequest["location"] = user.Location
	}
	if len(user.WebSite) > 0 {
		userRequest["webSite"] = user.WebSite
	}

	updtString := bson.M{
		"$set": userRequest,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
