package simulator

import (
	"github.com/stretchr/testify/mock"
)

func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

type CatchSimulatorMock struct {
	mock.Mock
}

func (c *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (ok bool) {
	args := c.Called(hunter, prey)
	return args.Bool(0)
}
