package mocks

import (
	"github.com/ildarusmanov/CommandResultor/interfaces"
	"github.com/stretchr/testify/mock"
)

type CommandResultMock struct {
	mock.Mock
}

func CreateCommandResultMock() *CommandResultMock {
	return new(CommandResultMock)
}

func (m *CommandResultMock) GetCode() int {
	args := m.Called()

	return args.Get(0).(int)
}

func (m *CommandResultMock) GetErrors() []error {
	args := m.Called()

	return args.Get(0).([]error)
}

func (m *CommandResultMock) GetValues() map[string]interfaces.CommandResultValue {
	args := m.Called()
	return args.Get(0).(map[string]interfaces.CommandResultValue)
}
