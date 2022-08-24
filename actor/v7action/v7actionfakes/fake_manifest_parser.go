// Code generated by counterfeiter. DO NOT EDIT.
package v7actionfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/v7action"
)

type FakeManifestParser struct {
	AppNamesStub        func() []string
	appNamesMutex       sync.RWMutex
	appNamesArgsForCall []struct {
	}
	appNamesReturns struct {
		result1 []string
	}
	appNamesReturnsOnCall map[int]struct {
		result1 []string
	}
	RawAppManifestStub        func(string) ([]byte, error)
	rawAppManifestMutex       sync.RWMutex
	rawAppManifestArgsForCall []struct {
		arg1 string
	}
	rawAppManifestReturns struct {
		result1 []byte
		result2 error
	}
	rawAppManifestReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestParser) AppNames() []string {
	fake.appNamesMutex.Lock()
	ret, specificReturn := fake.appNamesReturnsOnCall[len(fake.appNamesArgsForCall)]
	fake.appNamesArgsForCall = append(fake.appNamesArgsForCall, struct {
	}{})
	fake.recordInvocation("AppNames", []interface{}{})
	fake.appNamesMutex.Unlock()
	if fake.AppNamesStub != nil {
		return fake.AppNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.appNamesReturns
	return fakeReturns.result1
}

func (fake *FakeManifestParser) AppNamesCallCount() int {
	fake.appNamesMutex.RLock()
	defer fake.appNamesMutex.RUnlock()
	return len(fake.appNamesArgsForCall)
}

func (fake *FakeManifestParser) AppNamesCalls(stub func() []string) {
	fake.appNamesMutex.Lock()
	defer fake.appNamesMutex.Unlock()
	fake.AppNamesStub = stub
}

func (fake *FakeManifestParser) AppNamesReturns(result1 []string) {
	fake.appNamesMutex.Lock()
	defer fake.appNamesMutex.Unlock()
	fake.AppNamesStub = nil
	fake.appNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeManifestParser) AppNamesReturnsOnCall(i int, result1 []string) {
	fake.appNamesMutex.Lock()
	defer fake.appNamesMutex.Unlock()
	fake.AppNamesStub = nil
	if fake.appNamesReturnsOnCall == nil {
		fake.appNamesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.appNamesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeManifestParser) RawAppManifest(arg1 string) ([]byte, error) {
	fake.rawAppManifestMutex.Lock()
	ret, specificReturn := fake.rawAppManifestReturnsOnCall[len(fake.rawAppManifestArgsForCall)]
	fake.rawAppManifestArgsForCall = append(fake.rawAppManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RawAppManifest", []interface{}{arg1})
	fake.rawAppManifestMutex.Unlock()
	if fake.RawAppManifestStub != nil {
		return fake.RawAppManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.rawAppManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestParser) RawAppManifestCallCount() int {
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	return len(fake.rawAppManifestArgsForCall)
}

func (fake *FakeManifestParser) RawAppManifestCalls(stub func(string) ([]byte, error)) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = stub
}

func (fake *FakeManifestParser) RawAppManifestArgsForCall(i int) string {
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	argsForCall := fake.rawAppManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifestParser) RawAppManifestReturns(result1 []byte, result2 error) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = nil
	fake.rawAppManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) RawAppManifestReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.rawAppManifestMutex.Lock()
	defer fake.rawAppManifestMutex.Unlock()
	fake.RawAppManifestStub = nil
	if fake.rawAppManifestReturnsOnCall == nil {
		fake.rawAppManifestReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.rawAppManifestReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appNamesMutex.RLock()
	defer fake.appNamesMutex.RUnlock()
	fake.rawAppManifestMutex.RLock()
	defer fake.rawAppManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestParser) recordInvocation(key string, args []interface{}) {
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

var _ v7action.ManifestParser = new(FakeManifestParser)
