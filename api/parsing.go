package api

import (
	"io"
)

type Parser interface {
	LoadFromUrl(string) (io.Reader, error)
	LoadFromFile(string) (io.Reader, error)
	Match(io.Reader) bool
}
