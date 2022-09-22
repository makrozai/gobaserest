package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User es el modelo de usuario de la database
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name, omitempty" json:"name"`
	LastName    string             `bson:"lastName" json:"lastName,omitempty"`
	DateBirth   time.Time          `bson:"dateBirth" json:"dateBirth,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Description string             `bson:"description" json:"description,omitempty"`
	Location    string             `bson:"location" json:"location,omitempty"`
	WebSite     string             `bson:"webSite" json:"webSite,omitempty"`
}
