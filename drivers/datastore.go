package drivers

import (
	"github.com/jinzhu/gorm"
)

// DB - fetch instance of mongodb session
func DB() *gorm.DB {

	db, err := gorm.Open("postgres", "postgres://postgres@localhost/postgres?sslmode=disable")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	return db
}
