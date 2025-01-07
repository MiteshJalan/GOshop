// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "goshop/internal/order/dto"

	mock "github.com/stretchr/testify/mock"

	model "goshop/internal/order/model"

	paging "goshop/pkg/paging"
)

// IOrderRepository is an autogenerated mock type for the IOrderRepository type
type IOrderRepository struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, userID, lines
func (_m *IOrderRepository) CreateOrder(ctx context.Context, userID string, lines []*model.OrderLine) (*model.Order, error) {
	ret := _m.Called(ctx, userID, lines)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []*model.OrderLine) (*model.Order, error)); ok {
		return rf(ctx, userID, lines)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []*model.OrderLine) *model.Order); ok {
		r0 = rf(ctx, userID, lines)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []*model.OrderLine) error); ok {
		r1 = rf(ctx, userID, lines)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyOrders provides a mock function with given fields: ctx, req
func (_m *IOrderRepository) GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error) {
	ret := _m.Called(ctx, req)

	var r0 []*model.Order
	var r1 *paging.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListOrderReq) []*model.Order); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListOrderReq) *paging.Pagination); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paging.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *dto.ListOrderReq) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrderByID provides a mock function with given fields: ctx, id, preload
func (_m *IOrderRepository) GetOrderByID(ctx context.Context, id string, preload bool) (*model.Order, error) {
	ret := _m.Called(ctx, id, preload)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (*model.Order, error)); ok {
		return rf(ctx, id, preload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *model.Order); ok {
		r0 = rf(ctx, id, preload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, id, preload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrder provides a mock function with given fields: ctx, order
func (_m *IOrderRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIOrderRepository creates a new instance of IOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrderRepository {
	mock := &IOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
