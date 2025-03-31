package valueobject

import "github.com/google/uuid"

type Id struct {
	value string
}

func NewId() *Id {
	value, _ := uuid.NewV7()

	return &Id{
		value: value.String(),
	}
}

func NewValue(value string) *Id {
	return &Id{
		value: value,
	}
}

func (i *Id) Value() string {
	return i.value
}
