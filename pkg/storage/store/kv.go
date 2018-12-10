package store

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/twistedogic/spero/pkg/schema"
	bolt "go.etcd.io/bbolt"
)

type KV struct {
	*bolt.DB
}

func New(file string) (*KV, error) {
	db, err := bolt.Open(file, 0666, nil)
	return &KV{db}, err
}

func (k *KV) GetAll(id string) ([]schema.Match, error) {
	var output []schema.Match
	if err := k.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(id))
		if bucket == nil {
			return fmt.Errorf("bucket does not exist")
		}
		values := bucket.Get([]byte(id))
		if err := json.Unmarshal(values, &output); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return output, err
	}
	return output, nil
}

func (k *KV) GetLatest(id string) (schema.Match, error) {
	snapshots, err := k.GetAll(id)
	if err != nil || len(snapshots) == 0 {
		return schema.Match{}, err
	}
	sort.Sort(schema.ByLastUpdate{snapshots})
	return snapshots[0], nil
}

func (k *KV) Write(match schema.Match) error {
	id := match.MatchID
	input := []schema.Match{match}
	if err := k.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(id))
		if err != nil {
			return err
		}
		values, err := k.GetAll(match.MatchID)
		if err == nil {
			input = append(input, values...)
		}
		b, err := json.Marshal(input)
		if err != nil {
			return err
		}
		if err := bucket.Put([]byte(match.MatchID), b); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
