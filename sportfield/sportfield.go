package sportfield

import (
	"time"

	"github.com/ibadsatria/ucantplayalone/location"
)

// SportField entitiy
type SportField struct {
	ID          uint    `json:"id"`
	SportType   string  `json:"sport_type"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	HourlyPrice float64 `json:"hourly_price"`
	LocationID  uint
	Location    location.Location `json:"location"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   time.Time         `json:"deleted_at"`
}
