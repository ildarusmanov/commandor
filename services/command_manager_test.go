package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
// test constructor
func TestCreateCommandManager(t *testing.T) {
	m := CreateCommandManager(storage)

	assert.NotNil(t, m)
}

// test create command method
func TestCommandsManagerCreateCommand(t *testing.T) {
	m := CreateCommandManager(storage)
	cmd, err := m.CreateCommand(commandName, commandParams, commandOptions)

	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(cmd)
	assert.Equal(cmd.GetName(), commandName)
}

// test find command method
func TestCommandsManagerFindCommand(t *testing.T) {
	m := CreateCommandsManager(storage)

	cmd, err := m.FindCommand(commandName)

	// should finish without
	// error and return command
	assert := assert.New(t)
	assert.Nil(err)
	assert.NotNil(cmd)
	assert.Equal(cmd.GetName(), commandName)
}

// test drop command method
func TestCommandsManagerDropCommand(t *testing.T) {
	m := CreateCommandsManager(storage)

	errDrop := m.DropCommand(commandName)
	cmd, errFind := m.FindCommand(commandName)

	// command should be deleted withour error
	assert := assert.New(t)
	assert.Nil(errDrop)
	assert.NotNil(errFind)
	assert.Nil(cmd)
}
