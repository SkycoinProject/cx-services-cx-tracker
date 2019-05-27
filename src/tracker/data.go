package tracker

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/watercompany/cx-tracker/src/database/postgres"
)

type data interface {
	createOrUpdate(app *CxApplication) error
	getByHash(hash string) (CxApplication, error)
	getByGenesisHash(genesisHash string) (CxApplication, error)
	findAll() ([]CxApplication, error)
	updateServer(server *Server) error
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

func (s store) createOrUpdate(app *CxApplication) error {
	db := s.db.Begin()
	var dbError error
	for _, err := range db.Save(app).GetErrors() {
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

func (s store) getByHash(hash string) (CxApplication, error) {
	app := CxApplication{}
	record := s.db.Set("gorm:auto_preload", true).Find(&app, "hash = ?", hash)

	if record.RecordNotFound() {
		return app, errCannotFindApplication
	}

	if record.Error != nil {
		log.Errorf("Error occurred while fetching cx application by hash %v - %v", hash, record.Error)
		return app, errUnableToRead
	}

	return app, nil
}

func (s store) getByGenesisHash(genesisHash string) (CxApplication, error) {
	app := CxApplication{}
	record := s.db.Set("gorm:auto_preload", true).Find(&app, "config ->> 'genesisHash' = ?", genesisHash)

	if record.RecordNotFound() {
		return app, errCannotFindApplication
	}

	if record.Error != nil {
		log.Errorf("Error occurred while fetching cx application by genesis hash %v - %v", genesisHash, record.Error)
		return app, errUnableToRead
	}

	return app, nil
}

func (s store) findAll() ([]CxApplication, error) {
	apps := []CxApplication{}
	record := s.db.Set("gorm:auto_preload", true).Find(&apps)

	if errs := record.GetErrors(); len(errs) > 0 {
		for err := range errs {
			log.Error("Error occurred while fetching cx applications ", err)
		}
		return apps, errUnableToRead
	}
	return apps, nil
}

func (s store) updateServer(server *Server) error {
	db := s.db.Begin()
	var dbError error
	for _, err := range db.Save(server).GetErrors() {
		dbError = err
		log.Errorf("Error while updating server with address: %v in DB: %v", server.Address, err)
	}
	if dbError != nil {
		db.Rollback()
		return dbError
	}
	db.Commit()

	return nil
}
