package usecase

import (
	_model "github.com/ibadsatria/ucantplayalone/member"
	"github.com/ibadsatria/ucantplayalone/member/repository"
)

// MemberUsecase define contract method
type MemberUsecase interface {
	Store(m *_model.Member) (*_model.Member, error)
	Update(m *_model.Member) (*_model.Member, error)
	Delete(id int64) (bool, error)
	GetByID(id int64) *_model.Member
	GetByUsername(username string) (*_model.Member, error)
	GetByCity(city string) (*[]_model.Member, error)
}

type memberUcaseStruct struct {
	repoMember repository.MemberRepository
}

// NewMemberUsecase init new usecase
func NewMemberUsecase(r repository.MemberRepository) MemberUsecase {
	return &memberUcaseStruct{repoMember: r}
}

func (u *memberUcaseStruct) Store(m *_model.Member) (*_model.Member, error) {
	_, err := u.repoMember.Store(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (u *memberUcaseStruct) Update(m *_model.Member) (*_model.Member, error)       { return nil, nil }
func (u *memberUcaseStruct) Delete(id int64) (bool, error)                         { return false, nil }
func (u *memberUcaseStruct) GetByID(id int64) *_model.Member                       { return nil }
func (u *memberUcaseStruct) GetByUsername(username string) (*_model.Member, error) { return nil, nil }
func (u *memberUcaseStruct) GetByCity(city string) (*[]_model.Member, error)       { return nil, nil }
