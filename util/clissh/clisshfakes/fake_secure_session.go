// Code generated by counterfeiter. DO NOT EDIT.
package clisshfakes

import (
	"io"
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/util/clissh"
	"golang.org/x/crypto/ssh"
)

type FakeSecureSession struct {
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
	RequestPtyStub        func(string, int, int, ssh.TerminalModes) error
	requestPtyMutex       sync.RWMutex
	requestPtyArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
		arg4 ssh.TerminalModes
	}
	requestPtyReturns struct {
		result1 error
	}
	requestPtyReturnsOnCall map[int]struct {
		result1 error
	}
	SendRequestStub        func(string, bool, []byte) (bool, error)
	sendRequestMutex       sync.RWMutex
	sendRequestArgsForCall []struct {
		arg1 string
		arg2 bool
		arg3 []byte
	}
	sendRequestReturns struct {
		result1 bool
		result2 error
	}
	sendRequestReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ShellStub        func() error
	shellMutex       sync.RWMutex
	shellArgsForCall []struct {
	}
	shellReturns struct {
		result1 error
	}
	shellReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(string) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 string
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StderrPipeStub        func() (io.Reader, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct {
	}
	stderrPipeReturns struct {
		result1 io.Reader
		result2 error
	}
	stderrPipeReturnsOnCall map[int]struct {
		result1 io.Reader
		result2 error
	}
	StdinPipeStub        func() (io.WriteCloser, error)
	stdinPipeMutex       sync.RWMutex
	stdinPipeArgsForCall []struct {
	}
	stdinPipeReturns struct {
		result1 io.WriteCloser
		result2 error
	}
	stdinPipeReturnsOnCall map[int]struct {
		result1 io.WriteCloser
		result2 error
	}
	StdoutPipeStub        func() (io.Reader, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct {
	}
	stdoutPipeReturns struct {
		result1 io.Reader
		result2 error
	}
	stdoutPipeReturnsOnCall map[int]struct {
		result1 io.Reader
		result2 error
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

func (fake *FakeSecureSession) Close() error {
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

func (fake *FakeSecureSession) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureSession) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSecureSession) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) CloseReturnsOnCall(i int, result1 error) {
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

func (fake *FakeSecureSession) RequestPty(arg1 string, arg2 int, arg3 int, arg4 ssh.TerminalModes) error {
	fake.requestPtyMutex.Lock()
	ret, specificReturn := fake.requestPtyReturnsOnCall[len(fake.requestPtyArgsForCall)]
	fake.requestPtyArgsForCall = append(fake.requestPtyArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
		arg4 ssh.TerminalModes
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("RequestPty", []interface{}{arg1, arg2, arg3, arg4})
	fake.requestPtyMutex.Unlock()
	if fake.RequestPtyStub != nil {
		return fake.RequestPtyStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.requestPtyReturns
	return fakeReturns.result1
}

func (fake *FakeSecureSession) RequestPtyCallCount() int {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	return len(fake.requestPtyArgsForCall)
}

func (fake *FakeSecureSession) RequestPtyCalls(stub func(string, int, int, ssh.TerminalModes) error) {
	fake.requestPtyMutex.Lock()
	defer fake.requestPtyMutex.Unlock()
	fake.RequestPtyStub = stub
}

func (fake *FakeSecureSession) RequestPtyArgsForCall(i int) (string, int, int, ssh.TerminalModes) {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	argsForCall := fake.requestPtyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeSecureSession) RequestPtyReturns(result1 error) {
	fake.requestPtyMutex.Lock()
	defer fake.requestPtyMutex.Unlock()
	fake.RequestPtyStub = nil
	fake.requestPtyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) RequestPtyReturnsOnCall(i int, result1 error) {
	fake.requestPtyMutex.Lock()
	defer fake.requestPtyMutex.Unlock()
	fake.RequestPtyStub = nil
	if fake.requestPtyReturnsOnCall == nil {
		fake.requestPtyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.requestPtyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) SendRequest(arg1 string, arg2 bool, arg3 []byte) (bool, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.sendRequestMutex.Lock()
	ret, specificReturn := fake.sendRequestReturnsOnCall[len(fake.sendRequestArgsForCall)]
	fake.sendRequestArgsForCall = append(fake.sendRequestArgsForCall, struct {
		arg1 string
		arg2 bool
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("SendRequest", []interface{}{arg1, arg2, arg3Copy})
	fake.sendRequestMutex.Unlock()
	if fake.SendRequestStub != nil {
		return fake.SendRequestStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.sendRequestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureSession) SendRequestCallCount() int {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return len(fake.sendRequestArgsForCall)
}

func (fake *FakeSecureSession) SendRequestCalls(stub func(string, bool, []byte) (bool, error)) {
	fake.sendRequestMutex.Lock()
	defer fake.sendRequestMutex.Unlock()
	fake.SendRequestStub = stub
}

func (fake *FakeSecureSession) SendRequestArgsForCall(i int) (string, bool, []byte) {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	argsForCall := fake.sendRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecureSession) SendRequestReturns(result1 bool, result2 error) {
	fake.sendRequestMutex.Lock()
	defer fake.sendRequestMutex.Unlock()
	fake.SendRequestStub = nil
	fake.sendRequestReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) SendRequestReturnsOnCall(i int, result1 bool, result2 error) {
	fake.sendRequestMutex.Lock()
	defer fake.sendRequestMutex.Unlock()
	fake.SendRequestStub = nil
	if fake.sendRequestReturnsOnCall == nil {
		fake.sendRequestReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.sendRequestReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) Shell() error {
	fake.shellMutex.Lock()
	ret, specificReturn := fake.shellReturnsOnCall[len(fake.shellArgsForCall)]
	fake.shellArgsForCall = append(fake.shellArgsForCall, struct {
	}{})
	fake.recordInvocation("Shell", []interface{}{})
	fake.shellMutex.Unlock()
	if fake.ShellStub != nil {
		return fake.ShellStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.shellReturns
	return fakeReturns.result1
}

func (fake *FakeSecureSession) ShellCallCount() int {
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	return len(fake.shellArgsForCall)
}

func (fake *FakeSecureSession) ShellCalls(stub func() error) {
	fake.shellMutex.Lock()
	defer fake.shellMutex.Unlock()
	fake.ShellStub = stub
}

func (fake *FakeSecureSession) ShellReturns(result1 error) {
	fake.shellMutex.Lock()
	defer fake.shellMutex.Unlock()
	fake.ShellStub = nil
	fake.shellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) ShellReturnsOnCall(i int, result1 error) {
	fake.shellMutex.Lock()
	defer fake.shellMutex.Unlock()
	fake.ShellStub = nil
	if fake.shellReturnsOnCall == nil {
		fake.shellReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.shellReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) Start(arg1 string) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startReturns
	return fakeReturns.result1
}

func (fake *FakeSecureSession) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeSecureSession) StartCalls(stub func(string) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeSecureSession) StartArgsForCall(i int) string {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecureSession) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) StderrPipe() (io.Reader, error) {
	fake.stderrPipeMutex.Lock()
	ret, specificReturn := fake.stderrPipeReturnsOnCall[len(fake.stderrPipeArgsForCall)]
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct {
	}{})
	fake.recordInvocation("StderrPipe", []interface{}{})
	fake.stderrPipeMutex.Unlock()
	if fake.StderrPipeStub != nil {
		return fake.StderrPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stderrPipeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureSession) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

func (fake *FakeSecureSession) StderrPipeCalls(stub func() (io.Reader, error)) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = stub
}

func (fake *FakeSecureSession) StderrPipeReturns(result1 io.Reader, result2 error) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StderrPipeReturnsOnCall(i int, result1 io.Reader, result2 error) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = nil
	if fake.stderrPipeReturnsOnCall == nil {
		fake.stderrPipeReturnsOnCall = make(map[int]struct {
			result1 io.Reader
			result2 error
		})
	}
	fake.stderrPipeReturnsOnCall[i] = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdinPipe() (io.WriteCloser, error) {
	fake.stdinPipeMutex.Lock()
	ret, specificReturn := fake.stdinPipeReturnsOnCall[len(fake.stdinPipeArgsForCall)]
	fake.stdinPipeArgsForCall = append(fake.stdinPipeArgsForCall, struct {
	}{})
	fake.recordInvocation("StdinPipe", []interface{}{})
	fake.stdinPipeMutex.Unlock()
	if fake.StdinPipeStub != nil {
		return fake.StdinPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stdinPipeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureSession) StdinPipeCallCount() int {
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	return len(fake.stdinPipeArgsForCall)
}

func (fake *FakeSecureSession) StdinPipeCalls(stub func() (io.WriteCloser, error)) {
	fake.stdinPipeMutex.Lock()
	defer fake.stdinPipeMutex.Unlock()
	fake.StdinPipeStub = stub
}

func (fake *FakeSecureSession) StdinPipeReturns(result1 io.WriteCloser, result2 error) {
	fake.stdinPipeMutex.Lock()
	defer fake.stdinPipeMutex.Unlock()
	fake.StdinPipeStub = nil
	fake.stdinPipeReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdinPipeReturnsOnCall(i int, result1 io.WriteCloser, result2 error) {
	fake.stdinPipeMutex.Lock()
	defer fake.stdinPipeMutex.Unlock()
	fake.StdinPipeStub = nil
	if fake.stdinPipeReturnsOnCall == nil {
		fake.stdinPipeReturnsOnCall = make(map[int]struct {
			result1 io.WriteCloser
			result2 error
		})
	}
	fake.stdinPipeReturnsOnCall[i] = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdoutPipe() (io.Reader, error) {
	fake.stdoutPipeMutex.Lock()
	ret, specificReturn := fake.stdoutPipeReturnsOnCall[len(fake.stdoutPipeArgsForCall)]
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct {
	}{})
	fake.recordInvocation("StdoutPipe", []interface{}{})
	fake.stdoutPipeMutex.Unlock()
	if fake.StdoutPipeStub != nil {
		return fake.StdoutPipeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stdoutPipeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecureSession) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

func (fake *FakeSecureSession) StdoutPipeCalls(stub func() (io.Reader, error)) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = stub
}

func (fake *FakeSecureSession) StdoutPipeReturns(result1 io.Reader, result2 error) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) StdoutPipeReturnsOnCall(i int, result1 io.Reader, result2 error) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = nil
	if fake.stdoutPipeReturnsOnCall == nil {
		fake.stdoutPipeReturnsOnCall = make(map[int]struct {
			result1 io.Reader
			result2 error
		})
	}
	fake.stdoutPipeReturnsOnCall[i] = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureSession) Wait() error {
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

func (fake *FakeSecureSession) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureSession) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeSecureSession) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureSession) WaitReturnsOnCall(i int, result1 error) {
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

func (fake *FakeSecureSession) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecureSession) recordInvocation(key string, args []interface{}) {
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

var _ clissh.SecureSession = new(FakeSecureSession)
