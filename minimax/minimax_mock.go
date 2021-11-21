package minimax

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type StateMock struct {
	mock.Mock
}

func (m StateMock) Eval() int {
	args := m.Called()
	fmt.Println("eval ", args.Get(0))
	return args.Get(0).(int)
}

func (m StateMock) GetChildren(isMaximizer bool) []State {
	args := m.Called(isMaximizer)
	return args.Get(0).([]State)
}
