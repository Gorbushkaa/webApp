package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Request struct {
	Timestamp time.Time     `db:"timestamp"`
	Req_time  time.Duration `db:"req_time"`
	Url       string        `db:"Url"`
}

var db *gorm.DB
var err error

func InitialMigration() {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=database user=postgres password=simple dbname=postgres port=5432 sslmode=disable ",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Request{})
}

func GetRequests() []Request {
	var requests []Request
	db.Limit(30).Order("timestamp DESC").Find(&requests).Scan(&requests)
	return requests
}

func SetRequest(request Request) {
	db.Create(request)
}

func GetLastRequestByUrl(url string) Request {
	var request Request
	db.Where(fmt.Sprintf("url='%s'", url)).Last(&request).Scan(&request)
	return request
}
