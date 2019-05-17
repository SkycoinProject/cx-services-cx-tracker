package tracker

import (
	"encoding/json"
	"time"
)

// CxApplication represents DB record for CX Application run on separate chain
type CxApplication struct {
	ID        uint            `gorm:"primary_key" json:"id"`
	Hash      string          `json:"hash"`
	Config    json.RawMessage `json:"config"`
	ChainType chainType       `json:"chainType"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	DeletedAt *time.Time      `json:"-"`
	Servers   []Server        `json:"servers" gorm:"foreignkey:CxApplicationID;"`
}

// Server represents DB record for one node that reported to the service that's running one of existing CX Applications
type Server struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	Address         string     `json:"address"`
	CxApplicationID uint       `json:"-"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"-"`
}

type cxApplicationConfig struct {
	GenesisHash string `json:"genesisHash"`
}

type chainType string

const (
	cx chainType = "cx"
	// fiber chainType = "fiber" //TODO enable once integration with fiber is supported
)
