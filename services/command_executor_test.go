package services

import (
	"github.com/ildarusmanov/commandor/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

// test constructor
func TestCreateCommandExecutor(t *testing.T) {
	e := CreateCommandExecutor(commandsManager)

	assert.NotNil(t, e)
}

// test execute method
func TestCommandExecutorExecute(t *testing.T) {
	e := CreateCommandExecutor(commandsManager)

	resSync, errSync := e.Execute(false, cmd, params)
	resAsync, errAsync := e.Execucte(true, cmd, params)

	assert := assert.New(t)

	// sync command should finish without error
	assert.NotNil(resSync)
	assert.Nil(errSync)

	// async command should finish without error
	assert.NotNil(resAsync)
	assert.Nil(errAsync)
}
