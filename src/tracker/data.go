package tracker

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/watercompany/cx-tracker/src/database/postgres"
)

type store interface {
	create(app *CxApplication) error
	getBy(hash string) (CxApplication, error)
	findAll() ([]CxApplication, error)
}

type data struct {
	db *gorm.DB
}

func DefaultData() data {
	return NewData(postgres.DB)
}

func NewData(database *gorm.DB) data {
	return data{
		db: database,
	}
}

func (u data) create(app *CxApplication) error {
	db := u.db.Begin()
	var dbError error
	for _, err := range db.Create(app).GetErrors() {
		dbError = err
		log.Error("Error while creating new cx application in DB ", err)
	}
	if dbError != nil {
		db.Rollback()
		return dbError
	}
	db.Commit()

	return nil
}

func (u data) getBy(hash string) (app CxApplication, err error) {
	record := u.db.Set("gorm:auto_preload", true).Find(&app, "hash = ?", hash)

	if record.RecordNotFound() {
		err = errCannotFindUser
		return
	}
	if errs := record.GetErrors(); len(errs) > 0 {
		for err := range errs {
			log.Errorf("Error occurred while fetching cx application by hash %v - %v", hash, err)
		}
		err = errUnableToRead
		return
	}
	return
}

func (u data) findAll() (apps []CxApplication, err error) {
	record := u.db.Set("gorm:auto_preload", true).Find(&apps)

	if errs := record.GetErrors(); len(errs) > 0 {
		for err := range errs {
			log.Error("Error occurred while fetching cx applications ", err)
		}
		err = errUnableToRead
		return
	}
	return
}
