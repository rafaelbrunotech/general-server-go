package event

import "github.com/rafaelbrunoss/general-server-go/internal/common/domain/event"

type Publisher struct {
	observers []event.IObserver
}

func NewPublisher() (*Publisher, error) {
	return &Publisher{
		observers: []event.IObserver{},
	}, nil
}

func (p *Publisher) Publish(command event.ICommand) error {
	for _, observer := range p.observers {
		if observer.GetOperation() == command.GetOperation() {
			err := observer.Notify(command)

			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func (p *Publisher) Register(observer event.IObserver) error {
	p.observers = append(p.observers, observer)

	return nil
}
