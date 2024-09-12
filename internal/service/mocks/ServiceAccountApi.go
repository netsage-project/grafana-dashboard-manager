// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	models "github.com/grafana/grafana-openapi-client-go/models"
	mock "github.com/stretchr/testify/mock"

	types "github.com/esnet/gdg/internal/types"
)

// ServiceAccountApi is an autogenerated mock type for the ServiceAccountApi type
type ServiceAccountApi struct {
	mock.Mock
}

type ServiceAccountApi_Expecter struct {
	mock *mock.Mock
}

func (_m *ServiceAccountApi) EXPECT() *ServiceAccountApi_Expecter {
	return &ServiceAccountApi_Expecter{mock: &_m.Mock}
}

// CreateServiceAccount provides a mock function with given fields: name, role, expiration
func (_m *ServiceAccountApi) CreateServiceAccount(name string, role string, expiration int64) (*models.ServiceAccountDTO, error) {
	ret := _m.Called(name, role, expiration)

	if len(ret) == 0 {
		panic("no return value specified for CreateServiceAccount")
	}

	var r0 *models.ServiceAccountDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) (*models.ServiceAccountDTO, error)); ok {
		return rf(name, role, expiration)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) *models.ServiceAccountDTO); ok {
		r0 = rf(name, role, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ServiceAccountDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(name, role, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountApi_CreateServiceAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateServiceAccount'
type ServiceAccountApi_CreateServiceAccount_Call struct {
	*mock.Call
}

// CreateServiceAccount is a helper method to define mock.On call
//   - name string
//   - role string
//   - expiration int64
func (_e *ServiceAccountApi_Expecter) CreateServiceAccount(name interface{}, role interface{}, expiration interface{}) *ServiceAccountApi_CreateServiceAccount_Call {
	return &ServiceAccountApi_CreateServiceAccount_Call{Call: _e.mock.On("CreateServiceAccount", name, role, expiration)}
}

func (_c *ServiceAccountApi_CreateServiceAccount_Call) Run(run func(name string, role string, expiration int64)) *ServiceAccountApi_CreateServiceAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *ServiceAccountApi_CreateServiceAccount_Call) Return(_a0 *models.ServiceAccountDTO, _a1 error) *ServiceAccountApi_CreateServiceAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountApi_CreateServiceAccount_Call) RunAndReturn(run func(string, string, int64) (*models.ServiceAccountDTO, error)) *ServiceAccountApi_CreateServiceAccount_Call {
	_c.Call.Return(run)
	return _c
}

// CreateServiceAccountToken provides a mock function with given fields: name, role, expiration
func (_m *ServiceAccountApi) CreateServiceAccountToken(name int64, role string, expiration int64) (*models.NewAPIKeyResult, error) {
	ret := _m.Called(name, role, expiration)

	if len(ret) == 0 {
		panic("no return value specified for CreateServiceAccountToken")
	}

	var r0 *models.NewAPIKeyResult
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, int64) (*models.NewAPIKeyResult, error)); ok {
		return rf(name, role, expiration)
	}
	if rf, ok := ret.Get(0).(func(int64, string, int64) *models.NewAPIKeyResult); ok {
		r0 = rf(name, role, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.NewAPIKeyResult)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, string, int64) error); ok {
		r1 = rf(name, role, expiration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountApi_CreateServiceAccountToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateServiceAccountToken'
type ServiceAccountApi_CreateServiceAccountToken_Call struct {
	*mock.Call
}

// CreateServiceAccountToken is a helper method to define mock.On call
//   - name int64
//   - role string
//   - expiration int64
func (_e *ServiceAccountApi_Expecter) CreateServiceAccountToken(name interface{}, role interface{}, expiration interface{}) *ServiceAccountApi_CreateServiceAccountToken_Call {
	return &ServiceAccountApi_CreateServiceAccountToken_Call{Call: _e.mock.On("CreateServiceAccountToken", name, role, expiration)}
}

func (_c *ServiceAccountApi_CreateServiceAccountToken_Call) Run(run func(name int64, role string, expiration int64)) *ServiceAccountApi_CreateServiceAccountToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *ServiceAccountApi_CreateServiceAccountToken_Call) Return(_a0 *models.NewAPIKeyResult, _a1 error) *ServiceAccountApi_CreateServiceAccountToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountApi_CreateServiceAccountToken_Call) RunAndReturn(run func(int64, string, int64) (*models.NewAPIKeyResult, error)) *ServiceAccountApi_CreateServiceAccountToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAllServiceAccounts provides a mock function with given fields:
func (_m *ServiceAccountApi) DeleteAllServiceAccounts() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllServiceAccounts")
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

// ServiceAccountApi_DeleteAllServiceAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAllServiceAccounts'
type ServiceAccountApi_DeleteAllServiceAccounts_Call struct {
	*mock.Call
}

// DeleteAllServiceAccounts is a helper method to define mock.On call
func (_e *ServiceAccountApi_Expecter) DeleteAllServiceAccounts() *ServiceAccountApi_DeleteAllServiceAccounts_Call {
	return &ServiceAccountApi_DeleteAllServiceAccounts_Call{Call: _e.mock.On("DeleteAllServiceAccounts")}
}

func (_c *ServiceAccountApi_DeleteAllServiceAccounts_Call) Run(run func()) *ServiceAccountApi_DeleteAllServiceAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceAccountApi_DeleteAllServiceAccounts_Call) Return(_a0 []string) *ServiceAccountApi_DeleteAllServiceAccounts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceAccountApi_DeleteAllServiceAccounts_Call) RunAndReturn(run func() []string) *ServiceAccountApi_DeleteAllServiceAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteServiceAccountTokens provides a mock function with given fields: serviceId
func (_m *ServiceAccountApi) DeleteServiceAccountTokens(serviceId int64) []string {
	ret := _m.Called(serviceId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteServiceAccountTokens")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(int64) []string); ok {
		r0 = rf(serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ServiceAccountApi_DeleteServiceAccountTokens_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteServiceAccountTokens'
type ServiceAccountApi_DeleteServiceAccountTokens_Call struct {
	*mock.Call
}

// DeleteServiceAccountTokens is a helper method to define mock.On call
//   - serviceId int64
func (_e *ServiceAccountApi_Expecter) DeleteServiceAccountTokens(serviceId interface{}) *ServiceAccountApi_DeleteServiceAccountTokens_Call {
	return &ServiceAccountApi_DeleteServiceAccountTokens_Call{Call: _e.mock.On("DeleteServiceAccountTokens", serviceId)}
}

func (_c *ServiceAccountApi_DeleteServiceAccountTokens_Call) Run(run func(serviceId int64)) *ServiceAccountApi_DeleteServiceAccountTokens_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ServiceAccountApi_DeleteServiceAccountTokens_Call) Return(_a0 []string) *ServiceAccountApi_DeleteServiceAccountTokens_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceAccountApi_DeleteServiceAccountTokens_Call) RunAndReturn(run func(int64) []string) *ServiceAccountApi_DeleteServiceAccountTokens_Call {
	_c.Call.Return(run)
	return _c
}

// ListServiceAccounts provides a mock function with given fields:
func (_m *ServiceAccountApi) ListServiceAccounts() []*types.ServiceAccountDTOWithTokens {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListServiceAccounts")
	}

	var r0 []*types.ServiceAccountDTOWithTokens
	if rf, ok := ret.Get(0).(func() []*types.ServiceAccountDTOWithTokens); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.ServiceAccountDTOWithTokens)
		}
	}

	return r0
}

// ServiceAccountApi_ListServiceAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListServiceAccounts'
type ServiceAccountApi_ListServiceAccounts_Call struct {
	*mock.Call
}

// ListServiceAccounts is a helper method to define mock.On call
func (_e *ServiceAccountApi_Expecter) ListServiceAccounts() *ServiceAccountApi_ListServiceAccounts_Call {
	return &ServiceAccountApi_ListServiceAccounts_Call{Call: _e.mock.On("ListServiceAccounts")}
}

func (_c *ServiceAccountApi_ListServiceAccounts_Call) Run(run func()) *ServiceAccountApi_ListServiceAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ServiceAccountApi_ListServiceAccounts_Call) Return(_a0 []*types.ServiceAccountDTOWithTokens) *ServiceAccountApi_ListServiceAccounts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ServiceAccountApi_ListServiceAccounts_Call) RunAndReturn(run func() []*types.ServiceAccountDTOWithTokens) *ServiceAccountApi_ListServiceAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// ListServiceAccountsTokens provides a mock function with given fields: id
func (_m *ServiceAccountApi) ListServiceAccountsTokens(id int64) ([]*models.TokenDTO, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ListServiceAccountsTokens")
	}

	var r0 []*models.TokenDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]*models.TokenDTO, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) []*models.TokenDTO); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.TokenDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceAccountApi_ListServiceAccountsTokens_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListServiceAccountsTokens'
type ServiceAccountApi_ListServiceAccountsTokens_Call struct {
	*mock.Call
}

// ListServiceAccountsTokens is a helper method to define mock.On call
//   - id int64
func (_e *ServiceAccountApi_Expecter) ListServiceAccountsTokens(id interface{}) *ServiceAccountApi_ListServiceAccountsTokens_Call {
	return &ServiceAccountApi_ListServiceAccountsTokens_Call{Call: _e.mock.On("ListServiceAccountsTokens", id)}
}

func (_c *ServiceAccountApi_ListServiceAccountsTokens_Call) Run(run func(id int64)) *ServiceAccountApi_ListServiceAccountsTokens_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ServiceAccountApi_ListServiceAccountsTokens_Call) Return(_a0 []*models.TokenDTO, _a1 error) *ServiceAccountApi_ListServiceAccountsTokens_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ServiceAccountApi_ListServiceAccountsTokens_Call) RunAndReturn(run func(int64) ([]*models.TokenDTO, error)) *ServiceAccountApi_ListServiceAccountsTokens_Call {
	_c.Call.Return(run)
	return _c
}

// NewServiceAccountApi creates a new instance of ServiceAccountApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceAccountApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceAccountApi {
	mock := &ServiceAccountApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
