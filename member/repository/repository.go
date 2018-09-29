package repository

import (
	model "github.com/ibadsatria/ucantplayalone/member"
)

// MemberRepository repository of member
type MemberRepository interface {
	Store(m *model.Member) (uint, error)
	Update(m *model.Member) (*model.Member, error)
	Delete(id int64) (bool, error)
	GetByID(id int64) *model.Member
	GetByUsername(username string) (*model.Member, error)
	GetByCity(city string) (*[]model.Member, error)
}
