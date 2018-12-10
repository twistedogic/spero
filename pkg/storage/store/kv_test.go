package store

import (
	"os"
	"reflect"
	"testing"

	"github.com/twistedogic/spero/pkg/schema"
)

func TestNew(t *testing.T) {
	file := "temp.db"
	defer os.Remove(file)
	if _, err := New(file); err != nil {
		t.Error(err)
	}
}

func TestWrite(t *testing.T) {
	file := "temp.db"
	defer os.Remove(file)
	db, err := New(file)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	id := "test"
	input := schema.Match{MatchID: id}
	expect := []schema.Match{input}
	if err := db.Write(input); err != nil {
		t.Error(err)
	}
	out, err := db.GetAll(id)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(out, expect) {
		t.Fail()
	}
	latest, err := db.GetLatest(id)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(latest, input) {
		t.Fail()
	}
}
