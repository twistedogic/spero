package storage

import (
	"testing"

	"github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/gorm"
)

func TestCreate(t *testing.T) {
	defer testdb.Reset()
	type TestInput struct {
		ID      int `gorm:"PRIMARY_KEY"`
		Product string
		Price   float64
	}
	db, err := gorm.Open("testdb", "")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&TestInput{})

	db.Create(TestInput{ID: 1, Product: "test", Price: 1.02})
}
