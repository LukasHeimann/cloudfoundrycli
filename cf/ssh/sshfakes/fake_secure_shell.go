// Code generated by counterfeiter. DO NOT EDIT.
package sshfakes

import (
	"sync"

	sshCmd "github.com/LukasHeimann/cloudfoundrycli/v8/cf/ssh"
	"github.com/LukasHeimann/cloudfoundrycli/v8/cf/ssh/options"
)

type FakeSecureShell struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConnectStub        func(*options.SSHOptions) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		arg1 *options.SSHOptions
	}
	connectReturns struct {
		result1 error
	}
	connectReturnsOnCall map[int]struct {
		result1 error
	}
	InteractiveSessionStub        func() error
	interactiveSessionMutex       sync.RWMutex
	interactiveSessionArgsForCall []struct {
	}
	interactiveSessionReturns struct {
		result1 error
	}
	interactiveSessionReturnsOnCall map[int]struct {
		result1 error
	}
	LocalPortForwardStub        func() error
	localPortForwardMutex       sync.RWMutex
	localPortForwardArgsForCall []struct {
	}
	localPortForwardReturns struct {
		result1 error
	}
	localPortForwardReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecureShell) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShell) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureShell) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSecureShell) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) Connect(arg1 *options.SSHOptions) error {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		arg1 *options.SSHOptions
	}{arg1})
	fake.recordInvocation("Connect", []interface{}{arg1})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.connectReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShell) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeSecureShell) ConnectCalls(stub func(*options.SSHOptions) error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = stub
}

func (fake *FakeSecureShell) ConnectArgsForCall(i int) *options.SSHOptions {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	argsForCall := fake.connectArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecureShell) ConnectReturns(result1 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) ConnectReturnsOnCall(i int, result1 error) {
	fake.connectMutex.Lock()
	defer fake.connectMutex.Unlock()
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) InteractiveSession() error {
	fake.interactiveSessionMutex.Lock()
	ret, specificReturn := fake.interactiveSessionReturnsOnCall[len(fake.interactiveSessionArgsForCall)]
	fake.interactiveSessionArgsForCall = append(fake.interactiveSessionArgsForCall, struct {
	}{})
	fake.recordInvocation("InteractiveSession", []interface{}{})
	fake.interactiveSessionMutex.Unlock()
	if fake.InteractiveSessionStub != nil {
		return fake.InteractiveSessionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.interactiveSessionReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShell) InteractiveSessionCallCount() int {
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	return len(fake.interactiveSessionArgsForCall)
}

func (fake *FakeSecureShell) InteractiveSessionCalls(stub func() error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = stub
}

func (fake *FakeSecureShell) InteractiveSessionReturns(result1 error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = nil
	fake.interactiveSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) InteractiveSessionReturnsOnCall(i int, result1 error) {
	fake.interactiveSessionMutex.Lock()
	defer fake.interactiveSessionMutex.Unlock()
	fake.InteractiveSessionStub = nil
	if fake.interactiveSessionReturnsOnCall == nil {
		fake.interactiveSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.interactiveSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) LocalPortForward() error {
	fake.localPortForwardMutex.Lock()
	ret, specificReturn := fake.localPortForwardReturnsOnCall[len(fake.localPortForwardArgsForCall)]
	fake.localPortForwardArgsForCall = append(fake.localPortForwardArgsForCall, struct {
	}{})
	fake.recordInvocation("LocalPortForward", []interface{}{})
	fake.localPortForwardMutex.Unlock()
	if fake.LocalPortForwardStub != nil {
		return fake.LocalPortForwardStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.localPortForwardReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShell) LocalPortForwardCallCount() int {
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	return len(fake.localPortForwardArgsForCall)
}

func (fake *FakeSecureShell) LocalPortForwardCalls(stub func() error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = stub
}

func (fake *FakeSecureShell) LocalPortForwardReturns(result1 error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = nil
	fake.localPortForwardReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) LocalPortForwardReturnsOnCall(i int, result1 error) {
	fake.localPortForwardMutex.Lock()
	defer fake.localPortForwardMutex.Unlock()
	fake.LocalPortForwardStub = nil
	if fake.localPortForwardReturnsOnCall == nil {
		fake.localPortForwardReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.localPortForwardReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1
}

func (fake *FakeSecureShell) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureShell) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeSecureShell) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShell) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecureShell) recordInvocation(key string, args []interface{}) {
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

var _ sshCmd.SecureShell = new(FakeSecureShell)
