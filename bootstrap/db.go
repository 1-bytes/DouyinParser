package bootstrap

import (
	"DouyinParser/pkg/config"
	"DouyinParser/pkg/model"
	"gorm.io/gorm"
	"time"
)

// SetupDB used for init databases and ORM
func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	// set max db connections number
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// set max db idle connections number
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// expiration time for each connection
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")))
	// create and keep db structure
	migration(db)
}

// migration create and keep db structure.
func migration(db *gorm.DB) {
	_ = db.AutoMigrate()
}
