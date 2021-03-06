// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeDeleteOrganizationActor struct {
	DeleteOrganizationStub        func(orgName string) (v2action.Warnings, error)
	deleteOrganizationMutex       sync.RWMutex
	deleteOrganizationArgsForCall []struct {
		orgName string
	}
	deleteOrganizationReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganization(orgName string) (v2action.Warnings, error) {
	fake.deleteOrganizationMutex.Lock()
	fake.deleteOrganizationArgsForCall = append(fake.deleteOrganizationArgsForCall, struct {
		orgName string
	}{orgName})
	fake.recordInvocation("DeleteOrganization", []interface{}{orgName})
	fake.deleteOrganizationMutex.Unlock()
	if fake.DeleteOrganizationStub != nil {
		return fake.DeleteOrganizationStub(orgName)
	} else {
		return fake.deleteOrganizationReturns.result1, fake.deleteOrganizationReturns.result2
	}
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationCallCount() int {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return len(fake.deleteOrganizationArgsForCall)
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationArgsForCall(i int) string {
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return fake.deleteOrganizationArgsForCall[i].orgName
}

func (fake *FakeDeleteOrganizationActor) DeleteOrganizationReturns(result1 v2action.Warnings, result2 error) {
	fake.DeleteOrganizationStub = nil
	fake.deleteOrganizationReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteOrganizationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteOrganizationMutex.RLock()
	defer fake.deleteOrganizationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDeleteOrganizationActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.DeleteOrganizationActor = new(FakeDeleteOrganizationActor)
