package tracker

import (
	"encoding/json"
	"time"
)

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

type Server struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	Address         string     `json:"address"`
	CxApplicationID uint       `json:"-"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"-"`
}

type chainType string

const (
	cx    chainType = "cx"
	fiber chainType = "fiber"
)
