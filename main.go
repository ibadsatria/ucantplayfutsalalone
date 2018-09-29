package main

import (
	_ "database/sql"
	"fmt"
	"log"
	_ "net/url"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	cfg "github.com/ibadsatria/ucantplayalone/config/env"
	_ "github.com/ibadsatria/ucantplayalone/config/middleware"
	"github.com/ibadsatria/ucantplayalone/middleware"
	_ "github.com/labstack/echo"
	_ "github.com/lib/pq"

	_repoLoc "github.com/ibadsatria/ucantplayalone/location/repository"

	_httpMember "github.com/ibadsatria/ucantplayalone/member/delivery/http"
	_repoMember "github.com/ibadsatria/ucantplayalone/member/repository"
	_ucaseMember "github.com/ibadsatria/ucantplayalone/member/usecase"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUNs on DEBUG mode")
	}
}

func main() {
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)
	dbAddress := config.GetString(`database.host`)

	dbConn, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			dbAddress, dbUser, dbPass, dbName))
	defer dbConn.Close()

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// location setup
	repoLoc := _repoLoc.NewPgLocationRepo(dbConn)

	// member setup
	repoMember := _repoMember.NewPgMemberRepo(dbConn, repoLoc)
	ucaseMember := _ucaseMember.NewMemberUsecase(repoMember)
	_httpMember.NewMemberHTTPHandler(e, ucaseMember)

	// start the server
	e.Start(config.GetString("server.address"))
}
