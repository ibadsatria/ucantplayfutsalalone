package repository_test

import (
	"testing"
	"time"

	_repoLoc "github.com/ibadsatria/ucantplayalone/location/repository"
	_model "github.com/ibadsatria/ucantplayalone/schedule"
	_repoSch "github.com/ibadsatria/ucantplayalone/schedule/repository"

	"github.com/ibadsatria/ucantplayalone/playlevel"

	_modelMem "github.com/ibadsatria/ucantplayalone/member"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func TestInsertSchedule(t *testing.T) {

	db, err := gorm.Open("postgres", "host=localhost user=postgres password=qwerty123456 dbname=ucantplayalone sslmode=disable")
	if err != nil {
		t.Errorf("Errror %v", err)
	}
	defer db.Close()

	locRepo := _repoLoc.NewPgLocationRepo(db)
	schRepo := _repoSch.NewPgScheduleRepo(db, locRepo)

	var schedule _model.Schedule
	schedule = _model.Schedule{
		Level:        playlevel.FUN,
		Duration:     2,
		Time:         time.Now(),
		ReqPlayerNum: 12,
		LocationID:   1,
	}

	members := []_modelMem.Member{
		_modelMem.Member{Username: "ibad_satria"},
		_modelMem.Member{Username: "salmanm"},
		_modelMem.Member{Username: "username"},
	}

	schedule.Members = members

	schRepo.AddSchedule(&schedule)
}
