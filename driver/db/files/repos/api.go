package repos

import "os"

type provider struct {
	client *os.File
}

func NewProvider() *provider {
	return &provider{}
}