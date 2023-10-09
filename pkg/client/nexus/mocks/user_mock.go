// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	security "github.com/datadrivers/go-nexus-client/nexus3/schema/security"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *User) Get(id string) (*security.User, error) {
	ret := _m.Called(id)

	var r0 *security.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*security.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *security.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*security.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
