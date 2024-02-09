package arithmetic

import "errors"

type Arithmetic struct {
}

func NewAdapter() *Arithmetic {

	return &Arithmetic{}

}

func (arith Arithmetic) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arith Arithmetic) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arith Arithmetic) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arith Arithmetic) Division(a, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}
