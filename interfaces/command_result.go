package interfaces

type CommandResult interface {
    GetCode() int
    GetErrors() []error
    GetValues() map[string]CommandResultValue
}
