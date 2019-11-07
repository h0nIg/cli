// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSetOrgRoleActor struct {
	CreateOrgRoleByUserGUIDStub        func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)
	createOrgRoleByUserGUIDMutex       sync.RWMutex
	createOrgRoleByUserGUIDArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}
	createOrgRoleByUserGUIDReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createOrgRoleByUserGUIDReturnsOnCall map[int]struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	CreateOrgRoleByUserNameStub        func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)
	createOrgRoleByUserNameMutex       sync.RWMutex
	createOrgRoleByUserNameArgsForCall []struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}
	createOrgRoleByUserNameReturns struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	createOrgRoleByUserNameReturnsOnCall map[int]struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}
	GetOrganizationByNameStub        func(string) (v7action.Organization, v7action.Warnings, error)
	getOrganizationByNameMutex       sync.RWMutex
	getOrganizationByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationByNameReturns struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationByNameReturnsOnCall map[int]struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}
	GetUserStub        func(string, string) (v7action.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getUserReturns struct {
		result1 v7action.User
		result2 error
	}
	getUserReturnsOnCall map[int]struct {
		result1 v7action.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUID(arg1 constant.RoleType, arg2 string, arg3 string) (v7action.Role, v7action.Warnings, error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	ret, specificReturn := fake.createOrgRoleByUserGUIDReturnsOnCall[len(fake.createOrgRoleByUserGUIDArgsForCall)]
	fake.createOrgRoleByUserGUIDArgsForCall = append(fake.createOrgRoleByUserGUIDArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateOrgRoleByUserGUID", []interface{}{arg1, arg2, arg3})
	fake.createOrgRoleByUserGUIDMutex.Unlock()
	if fake.CreateOrgRoleByUserGUIDStub != nil {
		return fake.CreateOrgRoleByUserGUIDStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createOrgRoleByUserGUIDReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUIDCallCount() int {
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	return len(fake.createOrgRoleByUserGUIDArgsForCall)
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUIDCalls(stub func(constant.RoleType, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = stub
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUIDArgsForCall(i int) (constant.RoleType, string, string) {
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	argsForCall := fake.createOrgRoleByUserGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUIDReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = nil
	fake.createOrgRoleByUserGUIDReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserGUIDReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserGUIDMutex.Lock()
	defer fake.createOrgRoleByUserGUIDMutex.Unlock()
	fake.CreateOrgRoleByUserGUIDStub = nil
	if fake.createOrgRoleByUserGUIDReturnsOnCall == nil {
		fake.createOrgRoleByUserGUIDReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createOrgRoleByUserGUIDReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserName(arg1 constant.RoleType, arg2 string, arg3 string, arg4 string) (v7action.Role, v7action.Warnings, error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	ret, specificReturn := fake.createOrgRoleByUserNameReturnsOnCall[len(fake.createOrgRoleByUserNameArgsForCall)]
	fake.createOrgRoleByUserNameArgsForCall = append(fake.createOrgRoleByUserNameArgsForCall, struct {
		arg1 constant.RoleType
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateOrgRoleByUserName", []interface{}{arg1, arg2, arg3, arg4})
	fake.createOrgRoleByUserNameMutex.Unlock()
	if fake.CreateOrgRoleByUserNameStub != nil {
		return fake.CreateOrgRoleByUserNameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createOrgRoleByUserNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserNameCallCount() int {
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	return len(fake.createOrgRoleByUserNameArgsForCall)
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserNameCalls(stub func(constant.RoleType, string, string, string) (v7action.Role, v7action.Warnings, error)) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = stub
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserNameArgsForCall(i int) (constant.RoleType, string, string, string) {
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	argsForCall := fake.createOrgRoleByUserNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserNameReturns(result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = nil
	fake.createOrgRoleByUserNameReturns = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) CreateOrgRoleByUserNameReturnsOnCall(i int, result1 v7action.Role, result2 v7action.Warnings, result3 error) {
	fake.createOrgRoleByUserNameMutex.Lock()
	defer fake.createOrgRoleByUserNameMutex.Unlock()
	fake.CreateOrgRoleByUserNameStub = nil
	if fake.createOrgRoleByUserNameReturnsOnCall == nil {
		fake.createOrgRoleByUserNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Role
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createOrgRoleByUserNameReturnsOnCall[i] = struct {
		result1 v7action.Role
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByName(arg1 string) (v7action.Organization, v7action.Warnings, error) {
	fake.getOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationByNameReturnsOnCall[len(fake.getOrganizationByNameArgsForCall)]
	fake.getOrganizationByNameArgsForCall = append(fake.getOrganizationByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationByName", []interface{}{arg1})
	fake.getOrganizationByNameMutex.Unlock()
	if fake.GetOrganizationByNameStub != nil {
		return fake.GetOrganizationByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByNameCallCount() int {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	return len(fake.getOrganizationByNameArgsForCall)
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByNameCalls(stub func(string) (v7action.Organization, v7action.Warnings, error)) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = stub
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByNameArgsForCall(i int) string {
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByNameReturns(result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	fake.getOrganizationByNameReturns = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) GetOrganizationByNameReturnsOnCall(i int, result1 v7action.Organization, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationByNameMutex.Lock()
	defer fake.getOrganizationByNameMutex.Unlock()
	fake.GetOrganizationByNameStub = nil
	if fake.getOrganizationByNameReturnsOnCall == nil {
		fake.getOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Organization
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationByNameReturnsOnCall[i] = struct {
		result1 v7action.Organization
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSetOrgRoleActor) GetUser(arg1 string, arg2 string) (v7action.User, error) {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetUser", []interface{}{arg1, arg2})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSetOrgRoleActor) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeSetOrgRoleActor) GetUserCalls(stub func(string, string) (v7action.User, error)) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeSetOrgRoleActor) GetUserArgsForCall(i int) (string, string) {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	argsForCall := fake.getUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSetOrgRoleActor) GetUserReturns(result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeSetOrgRoleActor) GetUserReturnsOnCall(i int, result1 v7action.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 v7action.User
			result2 error
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 v7action.User
		result2 error
	}{result1, result2}
}

func (fake *FakeSetOrgRoleActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrgRoleByUserGUIDMutex.RLock()
	defer fake.createOrgRoleByUserGUIDMutex.RUnlock()
	fake.createOrgRoleByUserNameMutex.RLock()
	defer fake.createOrgRoleByUserNameMutex.RUnlock()
	fake.getOrganizationByNameMutex.RLock()
	defer fake.getOrganizationByNameMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSetOrgRoleActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v7.SetOrgRoleActor = new(FakeSetOrgRoleActor)
