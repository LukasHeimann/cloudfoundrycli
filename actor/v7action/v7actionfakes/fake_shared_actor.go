// Code generated by counterfeiter. DO NOT EDIT.
package v7actionfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/sharedaction"
	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7action"
)

type FakeSharedActor struct {
	GatherArchiveResourcesStub        func(string) ([]sharedaction.Resource, error)
	gatherArchiveResourcesMutex       sync.RWMutex
	gatherArchiveResourcesArgsForCall []struct {
		arg1 string
	}
	gatherArchiveResourcesReturns struct {
		result1 []sharedaction.Resource
		result2 error
	}
	gatherArchiveResourcesReturnsOnCall map[int]struct {
		result1 []sharedaction.Resource
		result2 error
	}
	GatherDirectoryResourcesStub        func(string) ([]sharedaction.Resource, error)
	gatherDirectoryResourcesMutex       sync.RWMutex
	gatherDirectoryResourcesArgsForCall []struct {
		arg1 string
	}
	gatherDirectoryResourcesReturns struct {
		result1 []sharedaction.Resource
		result2 error
	}
	gatherDirectoryResourcesReturnsOnCall map[int]struct {
		result1 []sharedaction.Resource
		result2 error
	}
	ZipArchiveResourcesStub        func(string, []sharedaction.Resource) (string, error)
	zipArchiveResourcesMutex       sync.RWMutex
	zipArchiveResourcesArgsForCall []struct {
		arg1 string
		arg2 []sharedaction.Resource
	}
	zipArchiveResourcesReturns struct {
		result1 string
		result2 error
	}
	zipArchiveResourcesReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ZipDirectoryResourcesStub        func(string, []sharedaction.Resource) (string, error)
	zipDirectoryResourcesMutex       sync.RWMutex
	zipDirectoryResourcesArgsForCall []struct {
		arg1 string
		arg2 []sharedaction.Resource
	}
	zipDirectoryResourcesReturns struct {
		result1 string
		result2 error
	}
	zipDirectoryResourcesReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSharedActor) GatherArchiveResources(arg1 string) ([]sharedaction.Resource, error) {
	fake.gatherArchiveResourcesMutex.Lock()
	ret, specificReturn := fake.gatherArchiveResourcesReturnsOnCall[len(fake.gatherArchiveResourcesArgsForCall)]
	fake.gatherArchiveResourcesArgsForCall = append(fake.gatherArchiveResourcesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GatherArchiveResources", []interface{}{arg1})
	fake.gatherArchiveResourcesMutex.Unlock()
	if fake.GatherArchiveResourcesStub != nil {
		return fake.GatherArchiveResourcesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.gatherArchiveResourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSharedActor) GatherArchiveResourcesCallCount() int {
	fake.gatherArchiveResourcesMutex.RLock()
	defer fake.gatherArchiveResourcesMutex.RUnlock()
	return len(fake.gatherArchiveResourcesArgsForCall)
}

func (fake *FakeSharedActor) GatherArchiveResourcesCalls(stub func(string) ([]sharedaction.Resource, error)) {
	fake.gatherArchiveResourcesMutex.Lock()
	defer fake.gatherArchiveResourcesMutex.Unlock()
	fake.GatherArchiveResourcesStub = stub
}

func (fake *FakeSharedActor) GatherArchiveResourcesArgsForCall(i int) string {
	fake.gatherArchiveResourcesMutex.RLock()
	defer fake.gatherArchiveResourcesMutex.RUnlock()
	argsForCall := fake.gatherArchiveResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSharedActor) GatherArchiveResourcesReturns(result1 []sharedaction.Resource, result2 error) {
	fake.gatherArchiveResourcesMutex.Lock()
	defer fake.gatherArchiveResourcesMutex.Unlock()
	fake.GatherArchiveResourcesStub = nil
	fake.gatherArchiveResourcesReturns = struct {
		result1 []sharedaction.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) GatherArchiveResourcesReturnsOnCall(i int, result1 []sharedaction.Resource, result2 error) {
	fake.gatherArchiveResourcesMutex.Lock()
	defer fake.gatherArchiveResourcesMutex.Unlock()
	fake.GatherArchiveResourcesStub = nil
	if fake.gatherArchiveResourcesReturnsOnCall == nil {
		fake.gatherArchiveResourcesReturnsOnCall = make(map[int]struct {
			result1 []sharedaction.Resource
			result2 error
		})
	}
	fake.gatherArchiveResourcesReturnsOnCall[i] = struct {
		result1 []sharedaction.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) GatherDirectoryResources(arg1 string) ([]sharedaction.Resource, error) {
	fake.gatherDirectoryResourcesMutex.Lock()
	ret, specificReturn := fake.gatherDirectoryResourcesReturnsOnCall[len(fake.gatherDirectoryResourcesArgsForCall)]
	fake.gatherDirectoryResourcesArgsForCall = append(fake.gatherDirectoryResourcesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GatherDirectoryResources", []interface{}{arg1})
	fake.gatherDirectoryResourcesMutex.Unlock()
	if fake.GatherDirectoryResourcesStub != nil {
		return fake.GatherDirectoryResourcesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.gatherDirectoryResourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSharedActor) GatherDirectoryResourcesCallCount() int {
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	return len(fake.gatherDirectoryResourcesArgsForCall)
}

func (fake *FakeSharedActor) GatherDirectoryResourcesCalls(stub func(string) ([]sharedaction.Resource, error)) {
	fake.gatherDirectoryResourcesMutex.Lock()
	defer fake.gatherDirectoryResourcesMutex.Unlock()
	fake.GatherDirectoryResourcesStub = stub
}

func (fake *FakeSharedActor) GatherDirectoryResourcesArgsForCall(i int) string {
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	argsForCall := fake.gatherDirectoryResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSharedActor) GatherDirectoryResourcesReturns(result1 []sharedaction.Resource, result2 error) {
	fake.gatherDirectoryResourcesMutex.Lock()
	defer fake.gatherDirectoryResourcesMutex.Unlock()
	fake.GatherDirectoryResourcesStub = nil
	fake.gatherDirectoryResourcesReturns = struct {
		result1 []sharedaction.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) GatherDirectoryResourcesReturnsOnCall(i int, result1 []sharedaction.Resource, result2 error) {
	fake.gatherDirectoryResourcesMutex.Lock()
	defer fake.gatherDirectoryResourcesMutex.Unlock()
	fake.GatherDirectoryResourcesStub = nil
	if fake.gatherDirectoryResourcesReturnsOnCall == nil {
		fake.gatherDirectoryResourcesReturnsOnCall = make(map[int]struct {
			result1 []sharedaction.Resource
			result2 error
		})
	}
	fake.gatherDirectoryResourcesReturnsOnCall[i] = struct {
		result1 []sharedaction.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) ZipArchiveResources(arg1 string, arg2 []sharedaction.Resource) (string, error) {
	var arg2Copy []sharedaction.Resource
	if arg2 != nil {
		arg2Copy = make([]sharedaction.Resource, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.zipArchiveResourcesMutex.Lock()
	ret, specificReturn := fake.zipArchiveResourcesReturnsOnCall[len(fake.zipArchiveResourcesArgsForCall)]
	fake.zipArchiveResourcesArgsForCall = append(fake.zipArchiveResourcesArgsForCall, struct {
		arg1 string
		arg2 []sharedaction.Resource
	}{arg1, arg2Copy})
	fake.recordInvocation("ZipArchiveResources", []interface{}{arg1, arg2Copy})
	fake.zipArchiveResourcesMutex.Unlock()
	if fake.ZipArchiveResourcesStub != nil {
		return fake.ZipArchiveResourcesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.zipArchiveResourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSharedActor) ZipArchiveResourcesCallCount() int {
	fake.zipArchiveResourcesMutex.RLock()
	defer fake.zipArchiveResourcesMutex.RUnlock()
	return len(fake.zipArchiveResourcesArgsForCall)
}

func (fake *FakeSharedActor) ZipArchiveResourcesCalls(stub func(string, []sharedaction.Resource) (string, error)) {
	fake.zipArchiveResourcesMutex.Lock()
	defer fake.zipArchiveResourcesMutex.Unlock()
	fake.ZipArchiveResourcesStub = stub
}

func (fake *FakeSharedActor) ZipArchiveResourcesArgsForCall(i int) (string, []sharedaction.Resource) {
	fake.zipArchiveResourcesMutex.RLock()
	defer fake.zipArchiveResourcesMutex.RUnlock()
	argsForCall := fake.zipArchiveResourcesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSharedActor) ZipArchiveResourcesReturns(result1 string, result2 error) {
	fake.zipArchiveResourcesMutex.Lock()
	defer fake.zipArchiveResourcesMutex.Unlock()
	fake.ZipArchiveResourcesStub = nil
	fake.zipArchiveResourcesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) ZipArchiveResourcesReturnsOnCall(i int, result1 string, result2 error) {
	fake.zipArchiveResourcesMutex.Lock()
	defer fake.zipArchiveResourcesMutex.Unlock()
	fake.ZipArchiveResourcesStub = nil
	if fake.zipArchiveResourcesReturnsOnCall == nil {
		fake.zipArchiveResourcesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.zipArchiveResourcesReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) ZipDirectoryResources(arg1 string, arg2 []sharedaction.Resource) (string, error) {
	var arg2Copy []sharedaction.Resource
	if arg2 != nil {
		arg2Copy = make([]sharedaction.Resource, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.zipDirectoryResourcesMutex.Lock()
	ret, specificReturn := fake.zipDirectoryResourcesReturnsOnCall[len(fake.zipDirectoryResourcesArgsForCall)]
	fake.zipDirectoryResourcesArgsForCall = append(fake.zipDirectoryResourcesArgsForCall, struct {
		arg1 string
		arg2 []sharedaction.Resource
	}{arg1, arg2Copy})
	fake.recordInvocation("ZipDirectoryResources", []interface{}{arg1, arg2Copy})
	fake.zipDirectoryResourcesMutex.Unlock()
	if fake.ZipDirectoryResourcesStub != nil {
		return fake.ZipDirectoryResourcesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.zipDirectoryResourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSharedActor) ZipDirectoryResourcesCallCount() int {
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	return len(fake.zipDirectoryResourcesArgsForCall)
}

func (fake *FakeSharedActor) ZipDirectoryResourcesCalls(stub func(string, []sharedaction.Resource) (string, error)) {
	fake.zipDirectoryResourcesMutex.Lock()
	defer fake.zipDirectoryResourcesMutex.Unlock()
	fake.ZipDirectoryResourcesStub = stub
}

func (fake *FakeSharedActor) ZipDirectoryResourcesArgsForCall(i int) (string, []sharedaction.Resource) {
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	argsForCall := fake.zipDirectoryResourcesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSharedActor) ZipDirectoryResourcesReturns(result1 string, result2 error) {
	fake.zipDirectoryResourcesMutex.Lock()
	defer fake.zipDirectoryResourcesMutex.Unlock()
	fake.ZipDirectoryResourcesStub = nil
	fake.zipDirectoryResourcesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) ZipDirectoryResourcesReturnsOnCall(i int, result1 string, result2 error) {
	fake.zipDirectoryResourcesMutex.Lock()
	defer fake.zipDirectoryResourcesMutex.Unlock()
	fake.ZipDirectoryResourcesStub = nil
	if fake.zipDirectoryResourcesReturnsOnCall == nil {
		fake.zipDirectoryResourcesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.zipDirectoryResourcesReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSharedActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gatherArchiveResourcesMutex.RLock()
	defer fake.gatherArchiveResourcesMutex.RUnlock()
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	fake.zipArchiveResourcesMutex.RLock()
	defer fake.zipArchiveResourcesMutex.RUnlock()
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSharedActor) recordInvocation(key string, args []interface{}) {
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

var _ v7action.SharedActor = new(FakeSharedActor)
