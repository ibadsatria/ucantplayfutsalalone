package repository

import (
	_model "github.com/ibadsatria/ucantplayalone/game"
	"github.com/jinzhu/gorm"
)

type pgGameRepository struct {
	DB *gorm.DB
}

// NewPgGameRepo init new repo
func NewPgGameRepo(db *gorm.DB) IGameRepository {
	return &pgGameRepository{DB: db}
}

func (r *pgGameRepository) AddGame(m *_model.Game) (bool, error) {
	err := r.DB.Create(m).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
