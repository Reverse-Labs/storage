package kv

import (
	"encoding/json"

	"github.com/Reverse-Labs/storage/pkg/storage"
)

type KV struct {
	Storage storage.Provider
}

func (kv *KV) Put(key string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return kv.Storage.WriteBytes(key, b)
}

func (kv *KV) Get(key string, v interface{}) error {
	b, err := kv.Storage.ReadBytes(key)

	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
}
