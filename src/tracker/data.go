package tracker

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/watercompany/cx-tracker/src/database/postgres"
)

type data interface {
	create(app *CxApplication) error
	getBy(hash string) (CxApplication, error)
	findAll() ([]CxApplication, error)
}

type store struct {
	db *gorm.DB
}

func defaultData() data {
	return newData(postgres.DB)
}

func newData(database *gorm.DB) data {
	return store{
		db: database,
	}
}

func (s store) create(app *CxApplication) error {
	db := s.db.Begin()
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

func (s store) getBy(hash string) (app CxApplication, err error) {
	record := s.db.Set("gorm:auto_preload", true).Find(&app, "hash = ?", hash)

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

func (s store) findAll() (apps []CxApplication, err error) {
	record := s.db.Set("gorm:auto_preload", true).Find(&apps)

	if errs := record.GetErrors(); len(errs) > 0 {
		for err := range errs {
			log.Error("Error occurred while fetching cx applications ", err)
		}
		err = errUnableToRead
		return
	}
	return
}
