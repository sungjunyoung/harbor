// Code generated by mockery v2.1.0. DO NOT EDIT.

package user

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/user/models"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) (models.Users, error) {
	ret := _m.Called(ctx, query)

	var r0 models.Users
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) models.Users); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}