package service

type ILogger interface {
	Error(error string)
	Info(text string)
	Warn(text string)
}
