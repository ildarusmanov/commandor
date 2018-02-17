package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test constructor
func TestCreateMemoryStorage(t *testing.T) {
	s, err := CreateMemoryStorage(options)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(s)
}

// test save in storage method
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
