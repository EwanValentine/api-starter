package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
	"time"
)

type ThingRepository struct {
	db *gorm.DB
}

func NewThingRepository(db *gorm.DB) *ThingRepository {
	return &ThingRepository{
		db,
	}
}

type NullTime struct {
	time.Time
	Valid bool
}

type Thing struct {
	ID        string `gorm:"primary_key:true"`
	Title     string `json:"title"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt NullTime
}

func (thing *ThingRepository) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func (repository *ThingRepository) FindAll() ([]Thing, error) {
	var things []Thing

	err := repository.db.Find(&things).Error

	if err != nil {
		return things, err
	}

	return things, nil
}
