package services

import (
	"github.com/ildarusmanov/commandor/tests/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCommandExecutor(t *testing.T) {
	e := CreateCommandExecutor(commandsManager)

	assert.NotNil(t, e)
}

func TestCommandExecutorExecute(t *testing.T) {
	e := CreateCommandExecutor(commandsManager)

	resSync, errSync := e.Execute(false, cmd, params)
	resAsync, errAsync := e.Execucte(true, cmd, params)

	assert := assert.New(t)
	assert.NotNil(resSync)
	assert.Nil(errSync)
	assert.NotNil(resAsync)
	assert.Nil(errAsync)
}
