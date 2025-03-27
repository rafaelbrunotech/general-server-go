package valueobject

type Money struct {
	value uint
}

func NewMoney() (*Money, error) {
	money := &Money{
		value: 0,
	}

	return money, nil
}

func (m *Money) Value() uint {
	return m.value
}
