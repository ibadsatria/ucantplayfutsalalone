package repository

import (
	"time"

	model "github.com/ibadsatria/ucantplayalone/location"
	"github.com/jinzhu/gorm"
)

type pgLocationRepository struct {
	DB *gorm.DB
}

// NewPgLocationRepo init new pg location repo
func NewPgLocationRepo(db *gorm.DB) LocationRepository {
	return &pgLocationRepository{DB: db}
}

func (r *pgLocationRepository) Store(l *model.Location) (uint, error) {
	now := time.Now()
	l.CreatedAt = now
	l.UpdatedAt = now

	err := r.DB.Create(l).Error

	if err != nil {
		return 0, err
	}

	return l.ID, nil
}
