// Code generated by mockery v2.23.2. DO NOT EDIT.

package service

import (
	domain "github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockACService is an autogenerated mock type for the ACService type
type MockACService struct {
	mock.Mock
}

type MockACService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockACService) EXPECT() *MockACService_Expecter {
	return &MockACService_Expecter{mock: &_m.Mock}
}

// GetInfo provides a mock function with given fields:
func (_m *MockACService) GetInfo() (*domain.AC, error) {
	ret := _m.Called()

	var r0 *domain.AC
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.AC, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.AC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AC)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockACService_GetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInfo'
type MockACService_GetInfo_Call struct {
	*mock.Call
}

// GetInfo is a helper method to define mock.On call
func (_e *MockACService_Expecter) GetInfo() *MockACService_GetInfo_Call {
	return &MockACService_GetInfo_Call{Call: _e.mock.On("GetInfo")}
}

func (_c *MockACService_GetInfo_Call) Run(run func()) *MockACService_GetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockACService_GetInfo_Call) Return(_a0 *domain.AC, _a1 error) *MockACService_GetInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockACService_GetInfo_Call) RunAndReturn(run func() (*domain.AC, error)) *MockACService_GetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ToggleEnabled provides a mock function with given fields:
func (_m *MockACService) ToggleEnabled() (*domain.AC, error) {
	ret := _m.Called()

	var r0 *domain.AC
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.AC, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.AC); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AC)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockACService_ToggleEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToggleEnabled'
type MockACService_ToggleEnabled_Call struct {
	*mock.Call
}

// ToggleEnabled is a helper method to define mock.On call
func (_e *MockACService_Expecter) ToggleEnabled() *MockACService_ToggleEnabled_Call {
	return &MockACService_ToggleEnabled_Call{Call: _e.mock.On("ToggleEnabled")}
}

func (_c *MockACService_ToggleEnabled_Call) Run(run func()) *MockACService_ToggleEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockACService_ToggleEnabled_Call) Return(_a0 *domain.AC, _a1 error) *MockACService_ToggleEnabled_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockACService_ToggleEnabled_Call) RunAndReturn(run func() (*domain.AC, error)) *MockACService_ToggleEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateACSettings provides a mock function with given fields: desiredTemp, desiredHumidity
func (_m *MockACService) UpdateACSettings(desiredTemp float32, desiredHumidity float32) (*domain.AC, error) {
	ret := _m.Called(desiredTemp, desiredHumidity)

	var r0 *domain.AC
	var r1 error
	if rf, ok := ret.Get(0).(func(float32, float32) (*domain.AC, error)); ok {
		return rf(desiredTemp, desiredHumidity)
	}
	if rf, ok := ret.Get(0).(func(float32, float32) *domain.AC); ok {
		r0 = rf(desiredTemp, desiredHumidity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AC)
		}
	}

	if rf, ok := ret.Get(1).(func(float32, float32) error); ok {
		r1 = rf(desiredTemp, desiredHumidity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockACService_UpdateACSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateACSettings'
type MockACService_UpdateACSettings_Call struct {
	*mock.Call
}

// UpdateACSettings is a helper method to define mock.On call
//   - desiredTemp float32
//   - desiredHumidity float32
func (_e *MockACService_Expecter) UpdateACSettings(desiredTemp interface{}, desiredHumidity interface{}) *MockACService_UpdateACSettings_Call {
	return &MockACService_UpdateACSettings_Call{Call: _e.mock.On("UpdateACSettings", desiredTemp, desiredHumidity)}
}

func (_c *MockACService_UpdateACSettings_Call) Run(run func(desiredTemp float32, desiredHumidity float32)) *MockACService_UpdateACSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float32), args[1].(float32))
	})
	return _c
}

func (_c *MockACService_UpdateACSettings_Call) Return(_a0 *domain.AC, _a1 error) *MockACService_UpdateACSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockACService_UpdateACSettings_Call) RunAndReturn(run func(float32, float32) (*domain.AC, error)) *MockACService_UpdateACSettings_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockACService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockACService creates a new instance of MockACService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockACService(t mockConstructorTestingTNewMockACService) *MockACService {
	mock := &MockACService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}