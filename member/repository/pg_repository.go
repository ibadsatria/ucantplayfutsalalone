package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	model "github.com/ibadsatria/ucantplayalone/member"

	_locRepo "github.com/ibadsatria/ucantplayalone/location/repository"
)

type pgMemberRepository struct {
	DB      *gorm.DB
	locRepo _locRepo.LocationRepository
}

// NewPgMemberRepo init new postgresql repo
func NewPgMemberRepo(db *gorm.DB, locationRepo _locRepo.LocationRepository) MemberRepository {
	return &pgMemberRepository{DB: db, locRepo: locationRepo}
}

func (r *pgMemberRepository) getByUsername(username string) (*model.Member, error) {
	var member model.Member
	err := r.DB.Where("username=?", username).
		First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *pgMemberRepository) Store(m *model.Member) (uint, error) {
	if _, err := r.getByUsername(m.Username); err == nil {
		log.Println(model.CONFLICT_ERROR)
		return 0, model.CONFLICT_ERROR
	}

	tx := r.DB.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	defer tx.Rollback()

	fmt.Println(m)

	locID, err := r.locRepo.Store(&m.Location)
	if err != nil {
		return 0, err
	}
	m.LocationID = locID
	err = r.DB.Save(m).Error
	if err != nil {
		log.Fatal("failed registering new member.")
		return 0, err
	}

	tx.Commit()

	return m.ID, nil
}

func (r *pgMemberRepository) Update(m *model.Member) (*model.Member, error)        { return nil, nil }
func (r *pgMemberRepository) Delete(id int64) (bool, error)                        { return false, nil }
func (r *pgMemberRepository) GetByID(id int64) *model.Member                       { return nil }
func (r *pgMemberRepository) GetByUsername(username string) (*model.Member, error) { return nil, nil }

func (r *pgMemberRepository) GetByCity(city string) (*[]model.Member, error) {
	var members []model.Member

	err := r.DB.Joins("left join locations on members.location_id=locations.id").
		Preload("Location").
		Where("locations.city=?", city).Find(&members).Error

	if err != nil {
		return nil, err
	}

	return &members, nil
}
