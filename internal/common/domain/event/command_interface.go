package event

type ICommand interface {
	GetOperation() string
	Execute() error
}
