package db

import (
	"fmt"
	"time"

	"github.com/aygumov-g/service-url-shortener-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

func New(cfg config.DBConfig) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s ",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SSLMode,
	)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Storage{db: db}, nil
}

func (s *Storage) Get() *gorm.DB {
	return s.db
}

func (s *Storage) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
