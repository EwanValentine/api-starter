package drivers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Datastore - fetch instance of mongodb session
func DB() *gorm.DB {

	db, err := gorm.Open("postgres", "postgres://postgres@auth-datastore/postgres?sslmode=disable")

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
