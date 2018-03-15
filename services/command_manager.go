package services

import (
	"github.com/ildarusmanov/commandor/interfaces"
)

type CommandManager struct {
	storage interfaces.Storage
}

func CreateCommandManager(storage interfaces.Storage) *CommandManager {
	return &CommandManager{storage}
}
