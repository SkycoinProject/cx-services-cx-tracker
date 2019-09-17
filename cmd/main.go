package main

import (
	"github.com/SkycoinPro/cx-services-cx-tracker/src/app"
	"github.com/SkycoinPro/cx-services-cx-tracker/src/config"
	"github.com/SkycoinPro/cx-services-cx-tracker/src/database/postgres"
	"github.com/SkycoinPro/cx-services-cx-tracker/src/tracker"
)

// @title Skywire CX Tracker API
// @version 1.0
// @description This is a Skywire CX Tracker service used for saving configurations

// @host localhost:8083
// @BasePath /api/v1
func main() {
	config.Init("tracker-config")

	tearDown := postgres.Init()
	defer tearDown()

	uc := tracker.DefaultController()
	app.NewServer(
		uc,
	).Run()
}
