package interfaces

// command interface
type Command interface {
    Validate(args map[string]string) bool
    Execute(args map[string]string) error
}
