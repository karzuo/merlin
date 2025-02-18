// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	models "github.com/gojek/merlin/models"
	mock "github.com/stretchr/testify/mock"
)

// EnvironmentService is an autogenerated mock type for the EnvironmentService type
type EnvironmentService struct {
	mock.Mock
}

type EnvironmentService_GetDefaultEnvironment struct {
	*mock.Call
}

func (_m EnvironmentService_GetDefaultEnvironment) Return(_a0 *models.Environment, _a1 error) *EnvironmentService_GetDefaultEnvironment {
	return &EnvironmentService_GetDefaultEnvironment{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *EnvironmentService) OnGetDefaultEnvironment() *EnvironmentService_GetDefaultEnvironment {
	c := _m.On("GetDefaultEnvironment")
	return &EnvironmentService_GetDefaultEnvironment{Call: c}
}

func (_m *EnvironmentService) OnGetDefaultEnvironmentMatch(matchers ...interface{}) *EnvironmentService_GetDefaultEnvironment {
	c := _m.On("GetDefaultEnvironment", matchers...)
	return &EnvironmentService_GetDefaultEnvironment{Call: c}
}

// GetDefaultEnvironment provides a mock function with given fields:
func (_m *EnvironmentService) GetDefaultEnvironment() (*models.Environment, error) {
	ret := _m.Called()

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func() *models.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
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

type EnvironmentService_GetDefaultPredictionJobEnvironment struct {
	*mock.Call
}

func (_m EnvironmentService_GetDefaultPredictionJobEnvironment) Return(_a0 *models.Environment, _a1 error) *EnvironmentService_GetDefaultPredictionJobEnvironment {
	return &EnvironmentService_GetDefaultPredictionJobEnvironment{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *EnvironmentService) OnGetDefaultPredictionJobEnvironment() *EnvironmentService_GetDefaultPredictionJobEnvironment {
	c := _m.On("GetDefaultPredictionJobEnvironment")
	return &EnvironmentService_GetDefaultPredictionJobEnvironment{Call: c}
}

func (_m *EnvironmentService) OnGetDefaultPredictionJobEnvironmentMatch(matchers ...interface{}) *EnvironmentService_GetDefaultPredictionJobEnvironment {
	c := _m.On("GetDefaultPredictionJobEnvironment", matchers...)
	return &EnvironmentService_GetDefaultPredictionJobEnvironment{Call: c}
}

// GetDefaultPredictionJobEnvironment provides a mock function with given fields:
func (_m *EnvironmentService) GetDefaultPredictionJobEnvironment() (*models.Environment, error) {
	ret := _m.Called()

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func() *models.Environment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
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

type EnvironmentService_GetEnvironment struct {
	*mock.Call
}

func (_m EnvironmentService_GetEnvironment) Return(_a0 *models.Environment, _a1 error) *EnvironmentService_GetEnvironment {
	return &EnvironmentService_GetEnvironment{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *EnvironmentService) OnGetEnvironment(name string) *EnvironmentService_GetEnvironment {
	c := _m.On("GetEnvironment", name)
	return &EnvironmentService_GetEnvironment{Call: c}
}

func (_m *EnvironmentService) OnGetEnvironmentMatch(matchers ...interface{}) *EnvironmentService_GetEnvironment {
	c := _m.On("GetEnvironment", matchers...)
	return &EnvironmentService_GetEnvironment{Call: c}
}

// GetEnvironment provides a mock function with given fields: name
func (_m *EnvironmentService) GetEnvironment(name string) (*models.Environment, error) {
	ret := _m.Called(name)

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func(string) *models.Environment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type EnvironmentService_ListEnvironments struct {
	*mock.Call
}

func (_m EnvironmentService_ListEnvironments) Return(_a0 []*models.Environment, _a1 error) *EnvironmentService_ListEnvironments {
	return &EnvironmentService_ListEnvironments{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *EnvironmentService) OnListEnvironments(name string) *EnvironmentService_ListEnvironments {
	c := _m.On("ListEnvironments", name)
	return &EnvironmentService_ListEnvironments{Call: c}
}

func (_m *EnvironmentService) OnListEnvironmentsMatch(matchers ...interface{}) *EnvironmentService_ListEnvironments {
	c := _m.On("ListEnvironments", matchers...)
	return &EnvironmentService_ListEnvironments{Call: c}
}

// ListEnvironments provides a mock function with given fields: name
func (_m *EnvironmentService) ListEnvironments(name string) ([]*models.Environment, error) {
	ret := _m.Called(name)

	var r0 []*models.Environment
	if rf, ok := ret.Get(0).(func(string) []*models.Environment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type EnvironmentService_Save struct {
	*mock.Call
}

func (_m EnvironmentService_Save) Return(_a0 *models.Environment, _a1 error) *EnvironmentService_Save {
	return &EnvironmentService_Save{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *EnvironmentService) OnSave(env *models.Environment) *EnvironmentService_Save {
	c := _m.On("Save", env)
	return &EnvironmentService_Save{Call: c}
}

func (_m *EnvironmentService) OnSaveMatch(matchers ...interface{}) *EnvironmentService_Save {
	c := _m.On("Save", matchers...)
	return &EnvironmentService_Save{Call: c}
}

// Save provides a mock function with given fields: env
func (_m *EnvironmentService) Save(env *models.Environment) (*models.Environment, error) {
	ret := _m.Called(env)

	var r0 *models.Environment
	if rf, ok := ret.Get(0).(func(*models.Environment) *models.Environment); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Environment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Environment) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
