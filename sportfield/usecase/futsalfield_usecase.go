package usecase

import (
	_futsalfieldRepo "github.com/ibadsatria/ucantplayalone/futsalfield/repository"
)

type FutsalfieldRepository interface {
	GetByHourlyPriceRange(low, high float64) ([]*model.Futsalfield, error)
	GetByName(key string) (*model.Futsalfield, error)
	GetById(id int64) (*model.Futsalfield, error)
	Update(m *model.Futsalfield) (*model.Futsalfield, error)
	Store(m *model.Futsalfield) (int64, error)
	Delete(id int64) (bool, error)
}

type futsalfieldUsecase struct {
	futsalfieldRepo _futsalfieldRepo.FutsalfieldRepository 
}

func (u * futsalfieldUsecase) Store(m *model.Futsalfield) (int64, error) {
	id, err := u.futsalfieldRepo.Store(m)
	if err != nil {
		log
		return nil, err
	}
}