package bd

import (
	"context"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet graba el tweet en la database
func InsertTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("tweet")

	tweetRegister := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := col.InsertOne(ctx, tweetRegister)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.Hex(), true, nil
}
