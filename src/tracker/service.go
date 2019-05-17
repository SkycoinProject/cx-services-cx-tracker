package tracker

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Service handles tracker service layer
type Service struct {
	db data
}

//DefaultService creates new instance of service
func DefaultService() Service {
	return Service{
		db: defaultData(),
	}
}

// NewService prepares new instance of Service
func NewService(appStore data) Service {
	return Service{
		db: appStore,
	}
}

func (us *Service) createCxApplication(config []byte, address string) error {
	h := sha256.New()
	if _, err := h.Write(config); err != nil {
		log.Error("Error writing data: ", err)
		return err
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))

	app, err := us.db.getByHash(hash)
	if err != nil && err == errCannotFindApplication {
		configJSON := json.RawMessage(string(config))
		server := Server{Address: address}

		app = CxApplication{
			Hash:      hash,
			Config:    configJSON,
			ChainType: cx,
			Servers:   []Server{server},
		}
	} else {
		exsitingServer := false
		for _, server := range app.Servers {
			if address == server.Address {
				exsitingServer = true
				break
			}
		}

		if !exsitingServer {
			server := Server{Address: address}
			app.Servers = append(app.Servers, server)
		}
	}

	if err := us.db.createOrUpdate(&app); err != nil {
		return err
	}

	return nil
}

func (us *Service) getApplicationByGenesisHash(genesisHash string) (CxApplication, error) {
	return us.db.getByGenesisHash(genesisHash)
}

func (us *Service) findAllApplications() ([]CxApplication, error) {
	return us.db.findAll()
}
