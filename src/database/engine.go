package database

import (
	"gitlab.com/putuyaza/antrian/src/config"
	"gitlab.com/putuyaza/antrian/src/core"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connection() {
	dbConfig := config.Database()
	if dbConfig.Engine == "mysql" {
		Instance = core.Mysql()
	}
}

// MigrationFunc type func
type MigrationFunc func(db *gorm.DB) error

// migration collections
var upMigrations []MigrationFunc
var downMigration []MigrationFunc

// MigrationRegister new register migration
func MigrationRegister(up MigrationFunc, down MigrationFunc) {
	upMigrations = append(upMigrations, up)
	downMigration = append(downMigration, down)
}

// MigrationUp all available migration
func MigrationUp() error {
	for i := 0; i < len(upMigrations); i++ {
		// run up migration
		err := upMigrations[i](Instance)
		if err != nil {
			return err
		}
	}
	return nil
}

// MigrationDown all available migration
func MigrationDown() error {
	for i := 0; i < len(downMigration); i++ {
		// run down migration
		err := downMigration[i](Instance)
		if err != nil {
			return err
		}
	}
	return nil
}
