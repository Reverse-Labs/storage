package json

import (
	"encoding/json"
)

type JSON struct{}

func (j JSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j JSON) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
