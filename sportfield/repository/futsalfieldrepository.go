package repository

import (
	model "github.com/ibadsatria/ucantplayalone/futsalfield"
)

type FutsalfieldRepository interface {
	GetByHourlyPriceRange(low, high float64) ([]*model.Futsalfield, error)
	GetByName(key string) (*model.Futsalfield, error)
	GetByID(id int64) (*model.Futsalfield, error)
	Update(m *model.Futsalfield) (*model.Futsalfield, error)
	Store(m *model.Futsalfield) (int64, error)
	Delete(id int64) (bool, error)
}
