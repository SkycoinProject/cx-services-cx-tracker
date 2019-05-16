package tracker

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Service handles tracker service layer
type Service struct {
	db store
}

//DefaultService creates new instance of service
func DefaultService() Service {
	return Service{
		db: DefaultData(),
	}
}

// NewService prepares new instance of Service
func NewService(appStore store) Service {
	return Service{
		db: appStore,
	}
}

func (us *Service) createCxApplication(config []byte, address string) (string, error) {
	h := sha256.New()
	if _, err := h.Write(config); err != nil {
		log.Error("Error writing data: ", err)
		return "", err
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))
	configJSON := json.RawMessage(string(config))

	server := Server{Address: address}

	app := CxApplication{
		Hash:      hash,
		Config:    configJSON,
		ChainType: cx,
		Servers:   []Server{server},
	}

	if err := us.db.create(&app); err != nil {
		return "", err
	}

	return hash, nil
}

func (us *Service) getApplicationBy(hash string) (CxApplication, error) {
	return us.db.getBy(hash)
}

func (us *Service) findAllApplications() ([]CxApplication, error) {
	return us.db.findAll()
}
