package interfaces

type CommandResultValue interface {
	GetName() string
	GetType() string
	GetValue() string
}
