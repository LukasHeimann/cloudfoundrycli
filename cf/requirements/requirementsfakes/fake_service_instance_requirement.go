// Code generated by counterfeiter. DO NOT EDIT.
package requirementsfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/models"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/requirements"
)

type FakeServiceInstanceRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GetServiceInstanceStub        func() models.ServiceInstance
	getServiceInstanceMutex       sync.RWMutex
	getServiceInstanceArgsForCall []struct {
	}
	getServiceInstanceReturns struct {
		result1 models.ServiceInstance
	}
	getServiceInstanceReturnsOnCall map[int]struct {
		result1 models.ServiceInstance
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceInstanceRequirement) Execute() error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
	}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeServiceInstanceRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeServiceInstanceRequirement) ExecuteCalls(stub func() error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeServiceInstanceRequirement) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceInstanceRequirement) ExecuteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeServiceInstanceRequirement) GetServiceInstance() models.ServiceInstance {
	fake.getServiceInstanceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceReturnsOnCall[len(fake.getServiceInstanceArgsForCall)]
	fake.getServiceInstanceArgsForCall = append(fake.getServiceInstanceArgsForCall, struct {
	}{})
	fake.recordInvocation("GetServiceInstance", []interface{}{})
	fake.getServiceInstanceMutex.Unlock()
	if fake.GetServiceInstanceStub != nil {
		return fake.GetServiceInstanceStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getServiceInstanceReturns
	return fakeReturns.result1
}

func (fake *FakeServiceInstanceRequirement) GetServiceInstanceCallCount() int {
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	return len(fake.getServiceInstanceArgsForCall)
}

func (fake *FakeServiceInstanceRequirement) GetServiceInstanceCalls(stub func() models.ServiceInstance) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = stub
}

func (fake *FakeServiceInstanceRequirement) GetServiceInstanceReturns(result1 models.ServiceInstance) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	fake.getServiceInstanceReturns = struct {
		result1 models.ServiceInstance
	}{result1}
}

func (fake *FakeServiceInstanceRequirement) GetServiceInstanceReturnsOnCall(i int, result1 models.ServiceInstance) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	if fake.getServiceInstanceReturnsOnCall == nil {
		fake.getServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 models.ServiceInstance
		})
	}
	fake.getServiceInstanceReturnsOnCall[i] = struct {
		result1 models.ServiceInstance
	}{result1}
}

func (fake *FakeServiceInstanceRequirement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceInstanceRequirement) recordInvocation(key string, args []interface{}) {
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

var _ requirements.ServiceInstanceRequirement = new(FakeServiceInstanceRequirement)
