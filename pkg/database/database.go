package database

import (
	"bb.commandscape.com/cmd/commandscape/global/pkg/log"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

const (
	minBackoff = 5 * time.Millisecond
	maxBackoff = 5 * time.Second
)

type DatabaseConfig struct {
	Name string
}

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}

func Connect(config DatabaseConfig, tries int) (err error) {
	var delay time.Duration
	for i := 0; i < tries; i++ {

		db, err = gorm.Open("sqlite3", config.Name)

		if err != nil {
			log.Infof("Database not ready... retrying after %.2f seconds: %s", delay.Seconds(), err)
			if delay == 0 {
				delay = minBackoff
			} else {
				delay *= 2
			}
			if delay > maxBackoff {
				delay = maxBackoff
			}
			time.Sleep(delay)
			continue
		}
		return nil
	}
	return errors.New("unable to connect to the database")
}

func Close() error {
	return db.Close()
}
