package storage

import (
	"github.com/ildarusmanov/commandor/interfaces"
    "sync"
    "errors"
)

const commandNotFound = errors.New("Command not found")

type MemoryStorage struct {
    sync.Mutex
	options  map[string]string
	commands map[string]interfaces.Command
}

func CreateMemoryStorage(options map[string]string) *MemoryStorage {
	return &MemoryStorage{options, make(map[string]interfaces.Command)}
}

func (s *MemoryStorage) SaveCommand(c interfaces.Command) (interfaces.Command, error) {
    s.Lock()
    s.commands[c.GetName()] = c
    s.Unlock()

    r, ok := s.commands[c.GetName()]

    if !ok {
        return nil, commandNotFound
    }

    return r, nil
}
