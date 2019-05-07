package tracker

import (
	"crypto/sha256"
	"fmt"
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

func (s *Service) saveConfig(data []byte) string {
	h := sha256.New()
	h.Write(data)
	hash := fmt.Sprintf("%x", h.Sum(nil))

	s.store[hash] = data
	return hash
}

func (s *Service) readConfig(hash string) []byte {
	return s.store[hash]
}
