package location

import "time"

// Location entity
type Location struct {
	ID        uint      `json:"id"`
	Longitude float64   `json:"long"`
	Latitude  float64   `json:"lat"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
