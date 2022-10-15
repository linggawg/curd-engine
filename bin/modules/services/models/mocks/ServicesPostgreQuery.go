// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "engine/bin/modules/services/models/domain"

	mock "github.com/stretchr/testify/mock"
)

// ServicesPostgreQuery is an autogenerated mock type for the ServicesPostgre type
type ServicesPostgreQuery struct {
	mock.Mock
}

// FindOneByQueryID provides a mock function with given fields: ctx, queryID
func (_m *ServicesPostgreQuery) FindOneByQueryID(ctx context.Context, queryID string) (*models.Services, error) {
	ret := _m.Called(ctx, queryID)

	var r0 *models.Services
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Services); ok {
		r0 = rf(ctx, queryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, queryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneByServiceUrlAndMethodAndQueryIsNull provides a mock function with given fields: ctx, serviceUrl, method
func (_m *ServicesPostgreQuery) FindOneByServiceUrlAndMethodAndQueryIsNull(ctx context.Context, serviceUrl string, method string) (*models.Services, error) {
	ret := _m.Called(ctx, serviceUrl, method)

	var r0 *models.Services
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Services); ok {
		r0 = rf(ctx, serviceUrl, method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Services)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, serviceUrl, method)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServicesPostgreQuery interface {
	mock.TestingT
	Cleanup(func())
}

// NewServicesPostgreQuery creates a new instance of ServicesPostgreQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServicesPostgreQuery(t mockConstructorTestingTNewServicesPostgreQuery) *ServicesPostgreQuery {
	mock := &ServicesPostgreQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}