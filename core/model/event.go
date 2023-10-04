package model

import "github.com/google/uuid"

type Event struct {
	Id uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
}
