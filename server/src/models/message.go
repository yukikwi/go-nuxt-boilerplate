package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageText string // The main content of the message
}
