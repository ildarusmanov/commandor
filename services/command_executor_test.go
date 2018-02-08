package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCommandExecutor(t *testing.T) {
	e := CreateCommandExecutor(commands)

	assert.NotNil(t, e)
}

func TestExecCommandExecutor(t *testing.T) {
	e := CreateCommandExecutor(commands)

	resSync, errSync := e.Exec(false, cmd1, cmd2, cmd3)
	resAsync, errAsync := e.Exec(true, cmd1, cmd2, cmd3)

	assert := assert.New(t)
	assert.NotNil(resSync)
	assert.Nil(errSync)
	assert.NotNil(resAsync)
	assert.Nil(errAsync)
}
