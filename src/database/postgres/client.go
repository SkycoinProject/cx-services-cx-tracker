package postgres

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DB global store variable
var DB *gorm.DB

// Init creates a connection to database
func Init() func() {
	var err error
	DB, err = gorm.Open("postgres", dBInfo())
	DB.LogMode(viper.GetBool("database.log-mode"))
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
	}
	log.Info("Database connected")

	driver, err := postgres.WithInstance(DB.DB(), &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		viper.GetString("database.migration-source"),
		viper.GetString("database.name"), driver)
	if err != nil {
		log.Fatalf("Error while preparing database migration %v", err)
	}

	if err := m.Up(); err != nil {
		if strings.Contains(err.Error(), "no change") {
			log.Info("Nothing to migrate")
		} else {
			log.Fatal("Unable to migrate to the latest db version", err)
		}
	}
	log.Info("Migration process finished")

	return func() {
		log.Info("Disconnecting database")
		DB.Close()
		log.Debug("Database disconnected")
	}
}

func dBInfo() string {
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	sslmode := viper.GetString("database.sslmode")
	database := viper.GetString("database.name")

	dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user,
		password,
		host,
		port,
		database,
		sslmode,
	)
	log.Debug("Prepared connection string for db ", dbInfo)

	return dbInfo
}
