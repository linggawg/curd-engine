// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "engine/bin/modules/roles/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// RolesPostgreQuery is an autogenerated mock type for the RolesPostgre type
type RolesPostgreQuery struct {
	mock.Mock
}

// FindOneByID provides a mock function with given fields: ctx, id
func (_m *RolesPostgreQuery) FindOneByID(ctx context.Context, id string) (*models.Roles, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Roles
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Roles); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Roles)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRolesPostgreQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewRolesPostgreQuery creates a new instance of RolesPostgreQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRolesPostgreQuery(t mockConstructorTestingTNewRolesPostgreQuery) *RolesPostgreQuery {
	mock := &RolesPostgreQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}