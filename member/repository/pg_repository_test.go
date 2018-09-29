package repository_test

import (
	"fmt"
	"testing"

	_repoLoc "github.com/ibadsatria/ucantplayalone/location/repository"
	_repoMember "github.com/ibadsatria/ucantplayalone/member/repository"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// func TestInsertMember(t *testing.T) {
// 	db := pg.Connect(&pg.Options{
// 		User:     "postgres",
// 		Password: "qwerty123456",
// 		Database: "ucantplayalone",
// 	})
// 	defer db.Close()

// 	locRepo := _repoLoc.NewPgLocationRepo(db)
// 	memberRepo := _repoMember.NewPgMemberRepo(db, locRepo)

// 	newMember := &member.Member{Email: "salman@gmail.com", Fullname: "Salman M Ibadurrahman", Karma: 0, Username: "ibad_satria"}
// 	location := &location.Location{
// 		Longitude: -34837439.3434,
// 		Latitude:  -78678678676,
// 		City:      "Bandung",
// 	}

// 	newMember.Location = location

// 	id, err := memberRepo.Store(newMember)
// 	if err != nil {
// 		t.Fatalf("Error %v", err)
// 	}
// 	t.Log("ID: ", id)
// }

func TestGetByCity(t *testing.T) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=qwerty123456 dbname=ucantplayalone sslmode=disable")
	if err != nil {
		t.Errorf("Errror %v", err)
	}
	defer db.Close()

	locRepo := _repoLoc.NewPgLocationRepo(db)
	memberRepo := _repoMember.NewPgMemberRepo(db, locRepo)

	members, err := memberRepo.GetByCity("Bandung")
	if err != nil {
		t.Fatalf("failed get members by city %v", err)
	}

	for _, m := range *members {
		fmt.Println(m.Username)
		fmt.Println(m.Location.Longitude)
		fmt.Println(m.Location.Latitude)
		fmt.Println(m.Location.City)
		fmt.Println("==============")
	}
}
