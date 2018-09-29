package game

import "time"

// Game entity
type Game struct {
	MemberID   uint      `gorm:"primary_key" json:"member_id"`
	ScheduleID uint      `gorm:"primary_key" json:"schedule_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
