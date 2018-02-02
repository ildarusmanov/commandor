package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateCommandManager(t *testing.T) {
  m := CreateCommandManager()

  assert.NotNil(t, m)
}

func TestCommandsManagerCreateCommand(t *testing.T) {
  m := CreateCommandManager()
  cmd, err := m.CreateCommand(commandName, commandParams, commandOptions)

  assert := assert.New(t)
  assert.Nil(err)
  assert.NotNil(cmd)
  assert.Equal(cmd.GetName(), commandName)
}

