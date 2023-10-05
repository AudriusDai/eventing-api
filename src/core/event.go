package core

import (
	"github.com/audriusdai/eventing-api/core/model"
	"github.com/audriusdai/eventing-api/db"
)

var CreateEvent = func(event model.Event) (model.Event, error) {
	if r := db.DB.Create(&event); r.Error != nil {
		return event, r.Error
	}
	return event, nil
}
