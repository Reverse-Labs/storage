package storage

import (
	"io"
)

type Node interface {
	Prefix() string
}

type Provider interface {
	List(string) []string
	WriteBytes(string, []byte) error
	Write(string, io.Reader) error
	Read(string, io.Writer) error
	ReadBytes(string) ([]byte, error)
}
