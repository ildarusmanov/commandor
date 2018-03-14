package interfaces

type Storage interface {
	SaveCommand(command Command) (Command, error)
	FindCommand(name string) (Command, error)
}
