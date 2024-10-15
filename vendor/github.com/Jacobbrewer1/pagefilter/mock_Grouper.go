// Code generated by mockery. DO NOT EDIT.

package pagefilter

import mock "github.com/stretchr/testify/mock"

// MockGrouper is an autogenerated mock type for the Grouper type
type MockGrouper struct {
	mock.Mock
}

// Group provides a mock function with given fields:
func (_m *MockGrouper) Group() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Group")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// NewMockGrouper creates a new instance of MockGrouper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGrouper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGrouper {
	mock := &MockGrouper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
