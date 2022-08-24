// Code generated by counterfeiter. DO NOT EDIT.
package uifakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/util/configv3"
	"github.com/LukasHeimann/cloudfoundrycli/v8/util/ui"
)

type FakeConfig struct {
	ColorEnabledStub        func() configv3.ColorSetting
	colorEnabledMutex       sync.RWMutex
	colorEnabledArgsForCall []struct {
	}
	colorEnabledReturns struct {
		result1 configv3.ColorSetting
	}
	colorEnabledReturnsOnCall map[int]struct {
		result1 configv3.ColorSetting
	}
	IsTTYStub        func() bool
	isTTYMutex       sync.RWMutex
	isTTYArgsForCall []struct {
	}
	isTTYReturns struct {
		result1 bool
	}
	isTTYReturnsOnCall map[int]struct {
		result1 bool
	}
	LocaleStub        func() string
	localeMutex       sync.RWMutex
	localeArgsForCall []struct {
	}
	localeReturns struct {
		result1 string
	}
	localeReturnsOnCall map[int]struct {
		result1 string
	}
	TerminalWidthStub        func() int
	terminalWidthMutex       sync.RWMutex
	terminalWidthArgsForCall []struct {
	}
	terminalWidthReturns struct {
		result1 int
	}
	terminalWidthReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) ColorEnabled() configv3.ColorSetting {
	fake.colorEnabledMutex.Lock()
	ret, specificReturn := fake.colorEnabledReturnsOnCall[len(fake.colorEnabledArgsForCall)]
	fake.colorEnabledArgsForCall = append(fake.colorEnabledArgsForCall, struct {
	}{})
	stub := fake.ColorEnabledStub
	fakeReturns := fake.colorEnabledReturns
	fake.recordInvocation("ColorEnabled", []interface{}{})
	fake.colorEnabledMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) ColorEnabledCallCount() int {
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	return len(fake.colorEnabledArgsForCall)
}

func (fake *FakeConfig) ColorEnabledCalls(stub func() configv3.ColorSetting) {
	fake.colorEnabledMutex.Lock()
	defer fake.colorEnabledMutex.Unlock()
	fake.ColorEnabledStub = stub
}

func (fake *FakeConfig) ColorEnabledReturns(result1 configv3.ColorSetting) {
	fake.colorEnabledMutex.Lock()
	defer fake.colorEnabledMutex.Unlock()
	fake.ColorEnabledStub = nil
	fake.colorEnabledReturns = struct {
		result1 configv3.ColorSetting
	}{result1}
}

func (fake *FakeConfig) ColorEnabledReturnsOnCall(i int, result1 configv3.ColorSetting) {
	fake.colorEnabledMutex.Lock()
	defer fake.colorEnabledMutex.Unlock()
	fake.ColorEnabledStub = nil
	if fake.colorEnabledReturnsOnCall == nil {
		fake.colorEnabledReturnsOnCall = make(map[int]struct {
			result1 configv3.ColorSetting
		})
	}
	fake.colorEnabledReturnsOnCall[i] = struct {
		result1 configv3.ColorSetting
	}{result1}
}

func (fake *FakeConfig) IsTTY() bool {
	fake.isTTYMutex.Lock()
	ret, specificReturn := fake.isTTYReturnsOnCall[len(fake.isTTYArgsForCall)]
	fake.isTTYArgsForCall = append(fake.isTTYArgsForCall, struct {
	}{})
	stub := fake.IsTTYStub
	fakeReturns := fake.isTTYReturns
	fake.recordInvocation("IsTTY", []interface{}{})
	fake.isTTYMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) IsTTYCallCount() int {
	fake.isTTYMutex.RLock()
	defer fake.isTTYMutex.RUnlock()
	return len(fake.isTTYArgsForCall)
}

func (fake *FakeConfig) IsTTYCalls(stub func() bool) {
	fake.isTTYMutex.Lock()
	defer fake.isTTYMutex.Unlock()
	fake.IsTTYStub = stub
}

func (fake *FakeConfig) IsTTYReturns(result1 bool) {
	fake.isTTYMutex.Lock()
	defer fake.isTTYMutex.Unlock()
	fake.IsTTYStub = nil
	fake.isTTYReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) IsTTYReturnsOnCall(i int, result1 bool) {
	fake.isTTYMutex.Lock()
	defer fake.isTTYMutex.Unlock()
	fake.IsTTYStub = nil
	if fake.isTTYReturnsOnCall == nil {
		fake.isTTYReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isTTYReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) Locale() string {
	fake.localeMutex.Lock()
	ret, specificReturn := fake.localeReturnsOnCall[len(fake.localeArgsForCall)]
	fake.localeArgsForCall = append(fake.localeArgsForCall, struct {
	}{})
	stub := fake.LocaleStub
	fakeReturns := fake.localeReturns
	fake.recordInvocation("Locale", []interface{}{})
	fake.localeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) LocaleCallCount() int {
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	return len(fake.localeArgsForCall)
}

func (fake *FakeConfig) LocaleCalls(stub func() string) {
	fake.localeMutex.Lock()
	defer fake.localeMutex.Unlock()
	fake.LocaleStub = stub
}

func (fake *FakeConfig) LocaleReturns(result1 string) {
	fake.localeMutex.Lock()
	defer fake.localeMutex.Unlock()
	fake.LocaleStub = nil
	fake.localeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) LocaleReturnsOnCall(i int, result1 string) {
	fake.localeMutex.Lock()
	defer fake.localeMutex.Unlock()
	fake.LocaleStub = nil
	if fake.localeReturnsOnCall == nil {
		fake.localeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.localeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) TerminalWidth() int {
	fake.terminalWidthMutex.Lock()
	ret, specificReturn := fake.terminalWidthReturnsOnCall[len(fake.terminalWidthArgsForCall)]
	fake.terminalWidthArgsForCall = append(fake.terminalWidthArgsForCall, struct {
	}{})
	stub := fake.TerminalWidthStub
	fakeReturns := fake.terminalWidthReturns
	fake.recordInvocation("TerminalWidth", []interface{}{})
	fake.terminalWidthMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfig) TerminalWidthCallCount() int {
	fake.terminalWidthMutex.RLock()
	defer fake.terminalWidthMutex.RUnlock()
	return len(fake.terminalWidthArgsForCall)
}

func (fake *FakeConfig) TerminalWidthCalls(stub func() int) {
	fake.terminalWidthMutex.Lock()
	defer fake.terminalWidthMutex.Unlock()
	fake.TerminalWidthStub = stub
}

func (fake *FakeConfig) TerminalWidthReturns(result1 int) {
	fake.terminalWidthMutex.Lock()
	defer fake.terminalWidthMutex.Unlock()
	fake.TerminalWidthStub = nil
	fake.terminalWidthReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeConfig) TerminalWidthReturnsOnCall(i int, result1 int) {
	fake.terminalWidthMutex.Lock()
	defer fake.terminalWidthMutex.Unlock()
	fake.TerminalWidthStub = nil
	if fake.terminalWidthReturnsOnCall == nil {
		fake.terminalWidthReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.terminalWidthReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	fake.isTTYMutex.RLock()
	defer fake.isTTYMutex.RUnlock()
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	fake.terminalWidthMutex.RLock()
	defer fake.terminalWidthMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ ui.Config = new(FakeConfig)
