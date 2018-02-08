package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCommandManager(t *testing.T) {
	m := CreateCommandManager(storage)

	assert.NotNil(t, m)
}

func TestCommandsManagerCreateCommand(t *testing.T) {
	m := CreateCommandManager(storage)
	cmd, err := m.CreateCommand(commandName, commandParams, commandOptions)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(cmd)
	assert.Equal(cmd.GetName(), commandName)
}

func TestCommandsManagerFindCommand(t *testing.T) {
	m := CreateCommandsManager(storage)

	cmd, err := m.FindCommand(commandName)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(cmd)
}

func TestCommandsManagerDropCommand(t *testing.T) {
	m := CreateCommandsManager(storage)

	errDrop := m.DropCommand(commandName)
	cmd, errFind := m.FindCommand(commandName)

	assert := assert.New(t)
	assert.Nil(errDrop)
	assert.NotNil(errFind)
	assert.Nil(cmd)
}
