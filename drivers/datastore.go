package drivers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - fetch instance of postgres session
func DB(user, pass, host, name string) *gorm.DB {

	var connection string

	// Connection string
	connection = "postgres://" + user + ":" + pass + "@" + host + "/" + name + "?sslmode=disable"

	db, err := gorm.Open("postgres", connection)

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()

	// We have to panic at this stage as api is unusable
	if err != nil {
		panic(err)
	}

	return db
}
