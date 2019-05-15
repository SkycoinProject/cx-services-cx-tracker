package tracker

import (
	"crypto/sha256"
	"fmt"

	log "github.com/sirupsen/logrus"
)

//Service handles tracker service layer
type Service struct {
	store map[string][]byte
}

//DefaultService creates new instance of service
func DefaultService() Service {
	return Service{
		store: make(map[string][]byte),
	}
}

func (s *Service) saveConfig(data []byte) (string, error) {
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		log.Error("Error writing data: ", err)
		return "", err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))

	s.store[hash] = data
	return hash, nil
}

func (s *Service) readConfig(hash string) []byte {
	return s.store[hash]
}
