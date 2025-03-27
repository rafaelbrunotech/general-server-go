package event

type IPublisher interface {
	Publish(command ICommand) error
	Register(observer IObserver) error
}
