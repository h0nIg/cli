// This file was generated by counterfeiter
package api

import (
	. "github.com/cloudfoundry/cli/cf/api"

	"sync"
)

type FakeAppSecurityGroup struct {
	createMutex       sync.RWMutex
	createArgsForCall []ApplicationSecurityGroupFields

	createReturns struct {
		result1 error
	}
}

func (fake *FakeAppSecurityGroup) Create(arg1 ApplicationSecurityGroupFields) error {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.createArgsForCall = append(fake.createArgsForCall, arg1)

	return fake.createReturns.result1
}

func (fake *FakeAppSecurityGroup) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeAppSecurityGroup) CreateArgsForCall(i int) ApplicationSecurityGroupFields {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i]
}

func (fake *FakeAppSecurityGroup) CreateReturns(result1 error) {
	fake.createReturns = struct {
		result1 error
	}{result1}
}

var _ AppSecurityGroup = new(FakeAppSecurityGroup)
