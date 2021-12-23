package kv

import (
	"sync"

	"github.com/Reverse-Labs/storage/pkg/encoder"
	"github.com/Reverse-Labs/storage/pkg/storage"
)

type KV struct {
	storage   storage.Provider
	marshaler encoder.Marshaler
	lock      sync.Mutex
}

func New(s storage.Provider, m encoder.Marshaler) KV {
	return KV{storage: s, marshaler: m}
}

func (kv *KV) Put(key string, v interface{}) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()

	b, err := kv.marshaler.Marshal(v)

	if err != nil {
		return err
	}

	return kv.storage.WriteBytes(key, b)
}

func (kv *KV) Get(key string, v interface{}) error {
	kv.lock.Lock()
	defer kv.lock.Unlock()

	b, err := kv.storage.ReadBytes(key)

	if err != nil {
		return err
	}

	return kv.marshaler.Unmarshal(b, v)
}
