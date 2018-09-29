package member

import (
	"time"

	"github.com/ibadsatria/ucantplayalone/location"
)

// Member entity
type Member struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Reputation int    `json:"reputation"`
	LocationID uint
	Location   location.Location `gorm:"auto_preload" json:"loc,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  time.Time         `json:"deleted_at"`
}
