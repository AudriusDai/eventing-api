package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name         string    `gorm:"type:varchar;size:256"`
	Date         time.Time
	Languages    []string
	VideoQuality []string
	AudioQuality []string
	Invitees     []string
	Description  string `gorm:"type:varchar;size:1024"`
}
