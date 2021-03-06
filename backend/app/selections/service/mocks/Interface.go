// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import selections "github.com/tomney/finalfour/backend/app/selections"

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Interface) Create(_a0 selections.Selections) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(selections.Selections) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *Interface) List() ([]selections.Selections, error) {
	ret := _m.Called()

	var r0 []selections.Selections
	if rf, ok := ret.Get(0).(func() []selections.Selections); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]selections.Selections)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
