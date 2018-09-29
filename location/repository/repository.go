package repository

import model "github.com/ibadsatria/ucantplayalone/location"

// LocationRepository interface
type LocationRepository interface {
	Store(l *model.Location) (uint, error)
}
