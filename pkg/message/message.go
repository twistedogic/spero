package message

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/Jeffail/benthos/v3/public/service"
)

const (
	ContentHashKey = "content-hash"
)

func NewContentHashMessage(b []byte) *service.Message {
	msg := service.NewMessage(b)
	h := sha1.Sum(b)
	hash := hex.EncodeToString(h[:])
	msg.MetaSet(ContentHashKey, hash)
	return msg
}

func GetContentHash(msg *service.Message) (string, bool) {
	return msg.MetaGet(ContentHashKey)
}
