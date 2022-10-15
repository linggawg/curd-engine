// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "engine/bin/modules/users/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// UsersPostgreCommand is an autogenerated mock type for the UsersPostgre type
type UsersPostgreCommand struct {
	mock.Mock
}

// InsertOne provides a mock function with given fields: ctx, users
func (_m *UsersPostgreCommand) InsertOne(ctx context.Context, users *models.Users) error {
	ret := _m.Called(ctx, users)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Users) error); ok {
		r0 = rf(ctx, users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUsersPostgreCommand interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersPostgreCommand creates a new instance of UsersPostgreCommand. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersPostgreCommand(t mockConstructorTestingTNewUsersPostgreCommand) *UsersPostgreCommand {
	mock := &UsersPostgreCommand{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
