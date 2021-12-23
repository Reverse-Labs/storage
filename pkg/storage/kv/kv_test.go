package kv

import (
	"testing"

	"github.com/Reverse-Labs/storage/pkg/encoder/json"
	"github.com/Reverse-Labs/storage/pkg/s3"
)

var (
	config = s3.ConnectionProfile{
		URL:       "play.min.io",
		AccessKey: "Q3AM3UQ867SPQQA43P2F",
		SecretKey: "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG",
	}
	bucket = "go-storage"
)

func setup(t *testing.T) KV {
	s3, err := s3.FromCfg(config, bucket)
	if err != nil {
		t.Fatal(err)
	}

	return New(s3, json.JSON{})
}

func TestKV_Put(t *testing.T) {
	type args struct {
		key string
		v   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test put",
			args:    args{key: "test", v: []string{"some", "data"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := setup(t)

			if err := kv.Put(tt.args.key, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("KV.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKV_Get(t *testing.T) {
	type args struct {
		key string
		v   []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test get",
			args:    args{key: "test", v: make([]string, 0)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := setup(t)
			TestKV_Put(t)
			if err := kv.Get(tt.args.key, &tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("KV.Put() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(tt.args.v) == 0 {
				t.Fatal("No Data Read")
			}
		})
	}
}
