package store

import (
	bolt "go.etcd.io/bbolt"
)

type Meta struct {
	key    []byte
	bucket [][]byte
}

type IndexBucket struct {
	*bolt.Bucket
}

func (i *IndexBucket) Search(term []byte) []byte {
	return term
}
