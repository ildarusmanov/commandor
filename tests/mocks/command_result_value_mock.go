package mocks

import (
	"github.com/ildarusmanov/commandor/interfaces"
	"github.com/stretchr/testify/mock"
)

type CommandResultValueMock struct {
	mock.Mock
}

func CreateCommandResultValueMock() *CommandResultValueMock {
	return new(CommandResultValueMock)
}

func (m *CommandResultValueMock) GetName() string {
	args := m.Called()

	return args.Get(0).(string)
}

func (m *CommandResultValueMock) GetType() string {
	args := m.Called()

	return args.Get(0).(string)
}

func (m *CommandResultValueMock) GetValue() string {
	args := m.Called()
	return args.Get(0).(string)
}
