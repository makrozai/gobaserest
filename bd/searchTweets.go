package bd

import (
	"context"
	"log"
	"time"

	"github.com/makrozai/gobaserest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchTweets lee los tweets de un perfil
func SearchTweets(ID string, page int64) ([]*models.ResponseTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitterbaserest")
	col := db.Collection("tweet")

	var result []*models.ResponseTweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	pointer, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for pointer.Next(context.TODO()) {
		var tweet models.ResponseTweets
		err := pointer.Decode(&tweet)
		if err != nil {
			return result, false
		}
		result = append(result, &tweet)
	}
	return result, true
}
