package storage

import (
	"github.com/twistedogic/spero/pkg/schema"
)

type Storage interface {
	GetLatest(string) (schema.Match, error)
	GetAll(string) ([]schema.Match, error)
	Write(schema.Match) error
}
