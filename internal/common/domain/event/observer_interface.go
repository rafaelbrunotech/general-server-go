package event

type IObserver interface {
	GetOperation() string
	Notify(command ICommand) error
}
