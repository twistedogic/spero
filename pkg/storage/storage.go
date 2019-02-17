package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/twistedogic/spero/pkg/errors"
	"github.com/twistedogic/spero/pkg/metric"
)

type Storage struct {
	*gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{db}
}

func (s *Storage) WriteToDB(v interface{}) error {
	db := s.Create(v)
	errType := errors.NO_ERROR
	metric.UpdateDBMetric(v, errors.NewError(errType, db.Error))
	if db.Error != nil {
		errType = errors.DB_ERROR
		return db.Error
	}
	return nil
}
