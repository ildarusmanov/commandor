package interfaces

// command interface
type Command interface {
	GetName() string
	GetParamsList() map[string]string
	Validate(args map[string]string) bool
	Execute(args map[string]string) CommandResult
}
