// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	models "github.com/grafana/grafana-openapi-client-go/models"
	mock "github.com/stretchr/testify/mock"
)

// organizationUserCrudApi is an autogenerated mock type for the organizationUserCrudApi type
type organizationUserCrudApi struct {
	mock.Mock
}

type organizationUserCrudApi_Expecter struct {
	mock *mock.Mock
}

func (_m *organizationUserCrudApi) EXPECT() *organizationUserCrudApi_Expecter {
	return &organizationUserCrudApi_Expecter{mock: &_m.Mock}
}

// AddUserToOrg provides a mock function with given fields: role, orgSlug, userId
func (_m *organizationUserCrudApi) AddUserToOrg(role string, orgSlug string, userId int64) error {
	ret := _m.Called(role, orgSlug, userId)

	if len(ret) == 0 {
		panic("no return value specified for AddUserToOrg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(role, orgSlug, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// organizationUserCrudApi_AddUserToOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserToOrg'
type organizationUserCrudApi_AddUserToOrg_Call struct {
	*mock.Call
}

// AddUserToOrg is a helper method to define mock.On call
//   - role string
//   - orgSlug string
//   - userId int64
func (_e *organizationUserCrudApi_Expecter) AddUserToOrg(role interface{}, orgSlug interface{}, userId interface{}) *organizationUserCrudApi_AddUserToOrg_Call {
	return &organizationUserCrudApi_AddUserToOrg_Call{Call: _e.mock.On("AddUserToOrg", role, orgSlug, userId)}
}

func (_c *organizationUserCrudApi_AddUserToOrg_Call) Run(run func(role string, orgSlug string, userId int64)) *organizationUserCrudApi_AddUserToOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *organizationUserCrudApi_AddUserToOrg_Call) Return(_a0 error) *organizationUserCrudApi_AddUserToOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *organizationUserCrudApi_AddUserToOrg_Call) RunAndReturn(run func(string, string, int64) error) *organizationUserCrudApi_AddUserToOrg_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserFromOrg provides a mock function with given fields: orgId, userId
func (_m *organizationUserCrudApi) DeleteUserFromOrg(orgId string, userId int64) error {
	ret := _m.Called(orgId, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserFromOrg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(orgId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// organizationUserCrudApi_DeleteUserFromOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserFromOrg'
type organizationUserCrudApi_DeleteUserFromOrg_Call struct {
	*mock.Call
}

// DeleteUserFromOrg is a helper method to define mock.On call
//   - orgId string
//   - userId int64
func (_e *organizationUserCrudApi_Expecter) DeleteUserFromOrg(orgId interface{}, userId interface{}) *organizationUserCrudApi_DeleteUserFromOrg_Call {
	return &organizationUserCrudApi_DeleteUserFromOrg_Call{Call: _e.mock.On("DeleteUserFromOrg", orgId, userId)}
}

func (_c *organizationUserCrudApi_DeleteUserFromOrg_Call) Run(run func(orgId string, userId int64)) *organizationUserCrudApi_DeleteUserFromOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int64))
	})
	return _c
}

func (_c *organizationUserCrudApi_DeleteUserFromOrg_Call) Return(_a0 error) *organizationUserCrudApi_DeleteUserFromOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *organizationUserCrudApi_DeleteUserFromOrg_Call) RunAndReturn(run func(string, int64) error) *organizationUserCrudApi_DeleteUserFromOrg_Call {
	_c.Call.Return(run)
	return _c
}

// ListOrgUsers provides a mock function with given fields: orgId
func (_m *organizationUserCrudApi) ListOrgUsers(orgId int64) []*models.OrgUserDTO {
	ret := _m.Called(orgId)

	if len(ret) == 0 {
		panic("no return value specified for ListOrgUsers")
	}

	var r0 []*models.OrgUserDTO
	if rf, ok := ret.Get(0).(func(int64) []*models.OrgUserDTO); ok {
		r0 = rf(orgId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.OrgUserDTO)
		}
	}

	return r0
}

// organizationUserCrudApi_ListOrgUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListOrgUsers'
type organizationUserCrudApi_ListOrgUsers_Call struct {
	*mock.Call
}

// ListOrgUsers is a helper method to define mock.On call
//   - orgId int64
func (_e *organizationUserCrudApi_Expecter) ListOrgUsers(orgId interface{}) *organizationUserCrudApi_ListOrgUsers_Call {
	return &organizationUserCrudApi_ListOrgUsers_Call{Call: _e.mock.On("ListOrgUsers", orgId)}
}

func (_c *organizationUserCrudApi_ListOrgUsers_Call) Run(run func(orgId int64)) *organizationUserCrudApi_ListOrgUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *organizationUserCrudApi_ListOrgUsers_Call) Return(_a0 []*models.OrgUserDTO) *organizationUserCrudApi_ListOrgUsers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *organizationUserCrudApi_ListOrgUsers_Call) RunAndReturn(run func(int64) []*models.OrgUserDTO) *organizationUserCrudApi_ListOrgUsers_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserInOrg provides a mock function with given fields: role, orgSlug, userId
func (_m *organizationUserCrudApi) UpdateUserInOrg(role string, orgSlug string, userId int64) error {
	ret := _m.Called(role, orgSlug, userId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserInOrg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(role, orgSlug, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// organizationUserCrudApi_UpdateUserInOrg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserInOrg'
type organizationUserCrudApi_UpdateUserInOrg_Call struct {
	*mock.Call
}

// UpdateUserInOrg is a helper method to define mock.On call
//   - role string
//   - orgSlug string
//   - userId int64
func (_e *organizationUserCrudApi_Expecter) UpdateUserInOrg(role interface{}, orgSlug interface{}, userId interface{}) *organizationUserCrudApi_UpdateUserInOrg_Call {
	return &organizationUserCrudApi_UpdateUserInOrg_Call{Call: _e.mock.On("UpdateUserInOrg", role, orgSlug, userId)}
}

func (_c *organizationUserCrudApi_UpdateUserInOrg_Call) Run(run func(role string, orgSlug string, userId int64)) *organizationUserCrudApi_UpdateUserInOrg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int64))
	})
	return _c
}

func (_c *organizationUserCrudApi_UpdateUserInOrg_Call) Return(_a0 error) *organizationUserCrudApi_UpdateUserInOrg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *organizationUserCrudApi_UpdateUserInOrg_Call) RunAndReturn(run func(string, string, int64) error) *organizationUserCrudApi_UpdateUserInOrg_Call {
	_c.Call.Return(run)
	return _c
}

// newOrganizationUserCrudApi creates a new instance of organizationUserCrudApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newOrganizationUserCrudApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *organizationUserCrudApi {
	mock := &organizationUserCrudApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
