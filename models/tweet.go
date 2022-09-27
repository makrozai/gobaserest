package models

import "time"

//  Tweet es el formato o estructura que tendra el tweet en la database
type Tweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:" message,omitempty"`
	Date    time.Time `bson:"date" json:" date,omitempty"`
}
