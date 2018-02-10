package mocks

import (
	"github.com/ildarusmanov/commandor/interfaces"
	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

func CreateCommandMock() *CommandMock {
	return new(CommandMock)
}

func (m *CommandMock) GetName() string {
	args := m.Called()

	return args.Get(0).(string)
}

func (m *CommandMock) GetParamsList() map[string]string {
	args := m.Called()

	return args.Get(0).(map[string]string)
}

func (m *CommandMock) Execute(argsMap map[string]string) interfaces.CommandResult {
	args := m.Called()
	return args.Get(0).(interfaces.CommandResult)
}

func (m *CommandMock) Validate(argsMap map[string]string) bool {
	args := m.Called()
	return args.Get(0).(bool)
}
