package schedule

import (
	"time"

	"github.com/ibadsatria/ucantplayalone/member"
	"github.com/ibadsatria/ucantplayalone/sportfield"
)

// Schedule entity
type Schedule struct {
	ID           uint      `json:"id"`
	SportType    string    `json:"sport_type"`
	Level        int8      `json:"level"`
	RosterTime   time.Time `json:"time"`
	Duration     int       `json:"duration"`
	CreatorID    uint
	Creator      member.Member `json:"creator"`
	ReqPlayerNum int           `json:"req_player_number"`
	SportFieldID uint
	SportField   sportfield.SportField
	Members      []*member.Member `gorm:"many2many:games;association_jointable_foreignkey:schedule_id"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	DeletedAt    time.Time        `json:"deleted_at"`
}
