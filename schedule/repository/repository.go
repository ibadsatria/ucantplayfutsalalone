package repository

import (
	"github.com/ibadsatria/ucantplayalone/member"
	_model "github.com/ibadsatria/ucantplayalone/schedule"
)

// ScheduleRepository interface for repository
type ScheduleRepository interface {
	AddSchedule(sch *_model.Schedule) (*_model.Schedule, error)
	InsertMemberToSchedule(member member.Member) (bool, error)
	RemoveMember(memberID uint) (bool, error)
}
