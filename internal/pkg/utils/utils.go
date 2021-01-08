package utils

import (
	"errors"
	"fmt"
	"go-template/conf"
	"go-template/internal/pkg/db/postgres"
	"time"

	"github.com/jinzhu/gorm"

	"go-template/internal/app/models"
)

var (
	DefaultColor    string
	UrlPushConsumer string
	// NowFunc ...
	NowFunc = time.Now
)

const (
	SlugLength = 5
)

func AutoMigration() (err error) {
	//utils.GetConstant()
	dbPublic, err := postgres.GetDatabase("default")
	if err != nil {
		return
	}
	if _, err = dbPublic.DB().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); err != nil {
		return fmt.Errorf("error while creating DB extension 'uuid-ossp': %s", err)
	}
	//migrate table
	t := dbPublic.AutoMigrate(&models.Demo{})
	return t.Error
}

func IsErrNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func GetConstant(conf conf.AppConfig) {
}
