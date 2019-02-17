package sportradar

import (
	"encoding/json"
)

type Container struct {
	Doc []struct {
		Data json.RawMessage `json:"data"`
	} `json:"doc"`
}

func (c Container) GetData() []json.RawMessage {
	data := make([]json.RawMessage, len(c.Doc))
	for i, doc := range c.Doc {
		data[i] = doc.Data
	}
	return data
}
