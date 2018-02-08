package storage

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateMemoryStorage(t *testing.T) {
  s, err := CreateMemoryStorage(options)

  assert := assert.New(t)
  assert.Nil(err)
  assert.NotNil(s)
}

func TestMemoryStorageSaveCommand(t *testing.T) {
  s, _ := CreateMemoryStorage(options)

  savedCommand, saveErr := s.SaveCommand(command)
  
  assert.Nil(saveErr)
  assert.NotNil(SaveCommand)
  assert.Equal(command.GetName(), savedCommand.GetName())
  assert.Equal(commnad.GetParamsList(), savedCommand().GetParamsList())
}

func TestMemoryStorageFindCommand(t *testing.T) {
  s, _ := CreateMemoryStorage(options)

  savedCommand, saveErr := s.SaveCommand(command)
  foundCommand, foundErr := s.FindCommand(savedCommand.GetName())

  assert.Nil(saveErr)
  assert.Nil(foundErr)
  assert.NotNil(foundCommand)
}

