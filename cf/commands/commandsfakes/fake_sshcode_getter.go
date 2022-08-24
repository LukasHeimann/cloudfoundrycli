// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commandregistry"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/commands"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/flags"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
)

type FakeSSHCodeGetter struct {
	ExecuteStub        func(flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func() (string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
	}
	getReturns struct {
		result1 string
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct {
	}
	metaDataReturns struct {
		result1 commandregistry.CommandMetadata
	}
	metaDataReturnsOnCall map[int]struct {
		result1 commandregistry.CommandMetadata
	}
	RequirementsStub        func(requirements.Factory, flags.FlagContext) ([]requirements.Requirement, error)
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		arg1 requirements.Factory
		arg2 flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
		result2 error
	}
	requirementsReturnsOnCall map[int]struct {
		result1 []requirements.Requirement
		result2 error
	}
	SetDependencyStub        func(commandregistry.Dependency, bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		arg1 commandregistry.Dependency
		arg2 bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	setDependencyReturnsOnCall map[int]struct {
		result1 commandregistry.Command
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSSHCodeGetter) Execute(arg1 flags.FlagContext) error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 flags.FlagContext
	}{arg1})
	fake.recordInvocation("Execute", []interface{}{arg1})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeSSHCodeGetter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeSSHCodeGetter) ExecuteCalls(stub func(flags.FlagContext) error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeSSHCodeGetter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSSHCodeGetter) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHCodeGetter) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHCodeGetter) Get() (string, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
	}{})
	fake.recordInvocation("Get", []interface{}{})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSSHCodeGetter) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSSHCodeGetter) GetCalls(stub func() (string, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSSHCodeGetter) GetReturns(result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHCodeGetter) GetReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHCodeGetter) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	ret, specificReturn := fake.metaDataReturnsOnCall[len(fake.metaDataArgsForCall)]
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct {
	}{})
	fake.recordInvocation("MetaData", []interface{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metaDataReturns
	return fakeReturns.result1
}

func (fake *FakeSSHCodeGetter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeSSHCodeGetter) MetaDataCalls(stub func() commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = stub
}

func (fake *FakeSSHCodeGetter) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeSSHCodeGetter) MetaDataReturnsOnCall(i int, result1 commandregistry.CommandMetadata) {
	fake.metaDataMutex.Lock()
	defer fake.metaDataMutex.Unlock()
	fake.MetaDataStub = nil
	if fake.metaDataReturnsOnCall == nil {
		fake.metaDataReturnsOnCall = make(map[int]struct {
			result1 commandregistry.CommandMetadata
		})
	}
	fake.metaDataReturnsOnCall[i] = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeSSHCodeGetter) Requirements(arg1 requirements.Factory, arg2 flags.FlagContext) ([]requirements.Requirement, error) {
	fake.requirementsMutex.Lock()
	ret, specificReturn := fake.requirementsReturnsOnCall[len(fake.requirementsArgsForCall)]
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		arg1 requirements.Factory
		arg2 flags.FlagContext
	}{arg1, arg2})
	fake.recordInvocation("Requirements", []interface{}{arg1, arg2})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.requirementsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSSHCodeGetter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeSSHCodeGetter) RequirementsCalls(stub func(requirements.Factory, flags.FlagContext) ([]requirements.Requirement, error)) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = stub
}

func (fake *FakeSSHCodeGetter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	argsForCall := fake.requirementsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSSHCodeGetter) RequirementsReturns(result1 []requirements.Requirement, result2 error) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHCodeGetter) RequirementsReturnsOnCall(i int, result1 []requirements.Requirement, result2 error) {
	fake.requirementsMutex.Lock()
	defer fake.requirementsMutex.Unlock()
	fake.RequirementsStub = nil
	if fake.requirementsReturnsOnCall == nil {
		fake.requirementsReturnsOnCall = make(map[int]struct {
			result1 []requirements.Requirement
			result2 error
		})
	}
	fake.requirementsReturnsOnCall[i] = struct {
		result1 []requirements.Requirement
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHCodeGetter) SetDependency(arg1 commandregistry.Dependency, arg2 bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	ret, specificReturn := fake.setDependencyReturnsOnCall[len(fake.setDependencyArgsForCall)]
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		arg1 commandregistry.Dependency
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("SetDependency", []interface{}{arg1, arg2})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setDependencyReturns
	return fakeReturns.result1
}

func (fake *FakeSSHCodeGetter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeSSHCodeGetter) SetDependencyCalls(stub func(commandregistry.Dependency, bool) commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = stub
}

func (fake *FakeSSHCodeGetter) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	argsForCall := fake.setDependencyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSSHCodeGetter) SetDependencyReturns(result1 commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeSSHCodeGetter) SetDependencyReturnsOnCall(i int, result1 commandregistry.Command) {
	fake.setDependencyMutex.Lock()
	defer fake.setDependencyMutex.Unlock()
	fake.SetDependencyStub = nil
	if fake.setDependencyReturnsOnCall == nil {
		fake.setDependencyReturnsOnCall = make(map[int]struct {
			result1 commandregistry.Command
		})
	}
	fake.setDependencyReturnsOnCall[i] = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeSSHCodeGetter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSSHCodeGetter) recordInvocation(key string, args []interface{}) {
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

var _ commands.SSHCodeGetter = new(FakeSSHCodeGetter)
