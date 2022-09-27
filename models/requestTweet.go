package models

// RequestTweet captura del body, el mensaje que nos llega
type RequestTweet struct {
	Message string `bson:"message" json:"message"`
}
