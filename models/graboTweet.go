package models

import "time"

// GraboTweet modelo para grabar los tweet
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitemty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitemty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitemty"`
}
