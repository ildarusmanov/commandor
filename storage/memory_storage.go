package storage

import (
	"github.com/ildarusmanov/commandor/interfaces"
)

type MemoryStorage struct {
	options  map[string]string
	commands map[string]interfaces.Command
}

func CreateMemoryStorage(options map[string]string) *MemoryStorage {
	return &MemoryStorage{options, make(map[string]interfaces.Command)}
}
