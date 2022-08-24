// Code generated by counterfeiter. DO NOT EDIT.
package routefakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands/route"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
)

type FakeCreator struct {
	CreateRouteStub        func(string, string, int, bool, models.DomainFields, models.SpaceFields) (models.Route, error)
	createRouteMutex       sync.RWMutex
	createRouteArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 bool
		arg5 models.DomainFields
		arg6 models.SpaceFields
	}
	createRouteReturns struct {
		result1 models.Route
		result2 error
	}
	createRouteReturnsOnCall map[int]struct {
		result1 models.Route
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreator) CreateRoute(arg1 string, arg2 string, arg3 int, arg4 bool, arg5 models.DomainFields, arg6 models.SpaceFields) (models.Route, error) {
	fake.createRouteMutex.Lock()
	ret, specificReturn := fake.createRouteReturnsOnCall[len(fake.createRouteArgsForCall)]
	fake.createRouteArgsForCall = append(fake.createRouteArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
		arg4 bool
		arg5 models.DomainFields
		arg6 models.SpaceFields
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("CreateRoute", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createRouteMutex.Unlock()
	if fake.CreateRouteStub != nil {
		return fake.CreateRouteStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRouteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreator) CreateRouteCallCount() int {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	return len(fake.createRouteArgsForCall)
}

func (fake *FakeCreator) CreateRouteCalls(stub func(string, string, int, bool, models.DomainFields, models.SpaceFields) (models.Route, error)) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = stub
}

func (fake *FakeCreator) CreateRouteArgsForCall(i int) (string, string, int, bool, models.DomainFields, models.SpaceFields) {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	argsForCall := fake.createRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeCreator) CreateRouteReturns(result1 models.Route, result2 error) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = nil
	fake.createRouteReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeCreator) CreateRouteReturnsOnCall(i int, result1 models.Route, result2 error) {
	fake.createRouteMutex.Lock()
	defer fake.createRouteMutex.Unlock()
	fake.CreateRouteStub = nil
	if fake.createRouteReturnsOnCall == nil {
		fake.createRouteReturnsOnCall = make(map[int]struct {
			result1 models.Route
			result2 error
		})
	}
	fake.createRouteReturnsOnCall[i] = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreator) recordInvocation(key string, args []interface{}) {
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

var _ route.Creator = new(FakeCreator)
