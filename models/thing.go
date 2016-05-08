package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

// ThingRepository - Repository object for `things`
type ThingRepository struct {
	db *gorm.DB
}

// NewThingRepository - Create a new instance of `ThingRepository` database instance injected
func NewThingRepository(db *gorm.DB) *ThingRepository {
	return &ThingRepository{
		db,
	}
}

// Thing - Thing model
type Thing struct {
	ID        string `gorm:"primary_key:true"`
	Title     string `json:"title"`
	Amount    int    `json:"amount"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (thing *Thing) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

// FindAll - Find all of the things
func (repository *ThingRepository) FindAll() ([]Thing, error) {
	var things []Thing

	err := repository.db.Find(&things).Error

	if err != nil {
		return things, err
	}

	return things, nil
}

// Insert - Create a thing
func (repository *ThingRepository) Insert(thing Thing) error {
	return repository.db.Create(&thing).Error
}
