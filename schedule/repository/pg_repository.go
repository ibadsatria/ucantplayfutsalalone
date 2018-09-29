package repository

import (
	"github.com/ibadsatria/ucantplayalone/member"
	_model "github.com/ibadsatria/ucantplayalone/schedule"
	"github.com/jinzhu/gorm"

	_repoLoc "github.com/ibadsatria/ucantplayalone/location/repository"
)

type pgScheduleRepository struct {
	DB      *gorm.DB
	repoLoc _repoLoc.LocationRepository
}

// NewPgScheduleRepo create new instance
func NewPgScheduleRepo(db *gorm.DB, repoLoc _repoLoc.LocationRepository) ScheduleRepository {
	return &pgScheduleRepository{
		DB:      db,
		repoLoc: repoLoc,
	}
}

func (r *pgScheduleRepository) AddSchedule(sch *_model.Schedule) (*_model.Schedule, error) {
	tx := r.DB.Begin()

	var err error
	defer func(err error) (*_model.Schedule, error) {
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		tx.Commit()
		return sch, nil
	}(err)

	err = r.DB.Save(sch).Error

	return sch, nil
}
func (r *pgScheduleRepository) InsertMemberToSchedule(member member.Member) (bool, error) {
	return false, nil
}
func (r *pgScheduleRepository) RemoveMember(memberID uint) (bool, error) { return false, nil }
