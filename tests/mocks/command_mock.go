package mocks

import (
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

func (m *CommandMock) Validate(args map[string]string) bool {
    args := m.Called()
    return args.Get(0).(bool)
}

func (m *CommandMock) Validate(args map[string]string) bool {
    args := m.Called()
    return args.Get(0).(bool)
}
