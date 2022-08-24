// Code generated by counterfeiter. DO NOT EDIT.
package netfakes

import (
	"net/http"
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/net"
)

type FakeHTTPClientInterface struct {
	DoStub        func(*http.Request) (*http.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	DumpRequestStub        func(*http.Request)
	dumpRequestMutex       sync.RWMutex
	dumpRequestArgsForCall []struct {
		arg1 *http.Request
	}
	DumpResponseStub        func(*http.Response)
	dumpResponseMutex       sync.RWMutex
	dumpResponseArgsForCall []struct {
		arg1 *http.Response
	}
	ExecuteCheckRedirectStub        func(*http.Request, []*http.Request) error
	executeCheckRedirectMutex       sync.RWMutex
	executeCheckRedirectArgsForCall []struct {
		arg1 *http.Request
		arg2 []*http.Request
	}
	executeCheckRedirectReturns struct {
		result1 error
	}
	executeCheckRedirectReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHTTPClientInterface) Do(arg1 *http.Request) (*http.Response, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("Do", []interface{}{arg1})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.doReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPClientInterface) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeHTTPClientInterface) DoCalls(stub func(*http.Request) (*http.Response, error)) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *FakeHTTPClientInterface) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPClientInterface) DoReturns(result1 *http.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPClientInterface) DoReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPClientInterface) DumpRequest(arg1 *http.Request) {
	fake.dumpRequestMutex.Lock()
	fake.dumpRequestArgsForCall = append(fake.dumpRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("DumpRequest", []interface{}{arg1})
	fake.dumpRequestMutex.Unlock()
	if fake.DumpRequestStub != nil {
		fake.DumpRequestStub(arg1)
	}
}

func (fake *FakeHTTPClientInterface) DumpRequestCallCount() int {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	return len(fake.dumpRequestArgsForCall)
}

func (fake *FakeHTTPClientInterface) DumpRequestCalls(stub func(*http.Request)) {
	fake.dumpRequestMutex.Lock()
	defer fake.dumpRequestMutex.Unlock()
	fake.DumpRequestStub = stub
}

func (fake *FakeHTTPClientInterface) DumpRequestArgsForCall(i int) *http.Request {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	argsForCall := fake.dumpRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPClientInterface) DumpResponse(arg1 *http.Response) {
	fake.dumpResponseMutex.Lock()
	fake.dumpResponseArgsForCall = append(fake.dumpResponseArgsForCall, struct {
		arg1 *http.Response
	}{arg1})
	fake.recordInvocation("DumpResponse", []interface{}{arg1})
	fake.dumpResponseMutex.Unlock()
	if fake.DumpResponseStub != nil {
		fake.DumpResponseStub(arg1)
	}
}

func (fake *FakeHTTPClientInterface) DumpResponseCallCount() int {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	return len(fake.dumpResponseArgsForCall)
}

func (fake *FakeHTTPClientInterface) DumpResponseCalls(stub func(*http.Response)) {
	fake.dumpResponseMutex.Lock()
	defer fake.dumpResponseMutex.Unlock()
	fake.DumpResponseStub = stub
}

func (fake *FakeHTTPClientInterface) DumpResponseArgsForCall(i int) *http.Response {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	argsForCall := fake.dumpResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirect(arg1 *http.Request, arg2 []*http.Request) error {
	var arg2Copy []*http.Request
	if arg2 != nil {
		arg2Copy = make([]*http.Request, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.executeCheckRedirectMutex.Lock()
	ret, specificReturn := fake.executeCheckRedirectReturnsOnCall[len(fake.executeCheckRedirectArgsForCall)]
	fake.executeCheckRedirectArgsForCall = append(fake.executeCheckRedirectArgsForCall, struct {
		arg1 *http.Request
		arg2 []*http.Request
	}{arg1, arg2Copy})
	fake.recordInvocation("ExecuteCheckRedirect", []interface{}{arg1, arg2Copy})
	fake.executeCheckRedirectMutex.Unlock()
	if fake.ExecuteCheckRedirectStub != nil {
		return fake.ExecuteCheckRedirectStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeCheckRedirectReturns
	return fakeReturns.result1
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectCallCount() int {
	fake.executeCheckRedirectMutex.RLock()
	defer fake.executeCheckRedirectMutex.RUnlock()
	return len(fake.executeCheckRedirectArgsForCall)
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectCalls(stub func(*http.Request, []*http.Request) error) {
	fake.executeCheckRedirectMutex.Lock()
	defer fake.executeCheckRedirectMutex.Unlock()
	fake.ExecuteCheckRedirectStub = stub
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectArgsForCall(i int) (*http.Request, []*http.Request) {
	fake.executeCheckRedirectMutex.RLock()
	defer fake.executeCheckRedirectMutex.RUnlock()
	argsForCall := fake.executeCheckRedirectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectReturns(result1 error) {
	fake.executeCheckRedirectMutex.Lock()
	defer fake.executeCheckRedirectMutex.Unlock()
	fake.ExecuteCheckRedirectStub = nil
	fake.executeCheckRedirectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPClientInterface) ExecuteCheckRedirectReturnsOnCall(i int, result1 error) {
	fake.executeCheckRedirectMutex.Lock()
	defer fake.executeCheckRedirectMutex.Unlock()
	fake.ExecuteCheckRedirectStub = nil
	if fake.executeCheckRedirectReturnsOnCall == nil {
		fake.executeCheckRedirectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeCheckRedirectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHTTPClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	fake.executeCheckRedirectMutex.RLock()
	defer fake.executeCheckRedirectMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHTTPClientInterface) recordInvocation(key string, args []interface{}) {
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

var _ net.HTTPClientInterface = new(FakeHTTPClientInterface)
