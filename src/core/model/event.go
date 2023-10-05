package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Event struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name         string    `gorm:"type:varchar;size:256"`
	Date         time.Time
	Languages    pq.StringArray `gorm:"type:text[]"`
	VideoQuality pq.StringArray `gorm:"type:text[]"`
	AudioQuality pq.StringArray `gorm:"type:text[]"`
	Invitees     pq.StringArray `gorm:"type:text[]"`
	Description  string         `gorm:"type:varchar;size:1024"`
}
