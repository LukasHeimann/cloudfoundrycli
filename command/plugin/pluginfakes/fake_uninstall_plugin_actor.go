// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/pluginaction"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/plugin"
)

type FakeUninstallPluginActor struct {
	UninstallPluginStub        func(pluginaction.PluginUninstaller, string) error
	uninstallPluginMutex       sync.RWMutex
	uninstallPluginArgsForCall []struct {
		arg1 pluginaction.PluginUninstaller
		arg2 string
	}
	uninstallPluginReturns struct {
		result1 error
	}
	uninstallPluginReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUninstallPluginActor) UninstallPlugin(arg1 pluginaction.PluginUninstaller, arg2 string) error {
	fake.uninstallPluginMutex.Lock()
	ret, specificReturn := fake.uninstallPluginReturnsOnCall[len(fake.uninstallPluginArgsForCall)]
	fake.uninstallPluginArgsForCall = append(fake.uninstallPluginArgsForCall, struct {
		arg1 pluginaction.PluginUninstaller
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UninstallPlugin", []interface{}{arg1, arg2})
	fake.uninstallPluginMutex.Unlock()
	if fake.UninstallPluginStub != nil {
		return fake.UninstallPluginStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uninstallPluginReturns
	return fakeReturns.result1
}

func (fake *FakeUninstallPluginActor) UninstallPluginCallCount() int {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	return len(fake.uninstallPluginArgsForCall)
}

func (fake *FakeUninstallPluginActor) UninstallPluginCalls(stub func(pluginaction.PluginUninstaller, string) error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = stub
}

func (fake *FakeUninstallPluginActor) UninstallPluginArgsForCall(i int) (pluginaction.PluginUninstaller, string) {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	argsForCall := fake.uninstallPluginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUninstallPluginActor) UninstallPluginReturns(result1 error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = nil
	fake.uninstallPluginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUninstallPluginActor) UninstallPluginReturnsOnCall(i int, result1 error) {
	fake.uninstallPluginMutex.Lock()
	defer fake.uninstallPluginMutex.Unlock()
	fake.UninstallPluginStub = nil
	if fake.uninstallPluginReturnsOnCall == nil {
		fake.uninstallPluginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uninstallPluginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUninstallPluginActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUninstallPluginActor) recordInvocation(key string, args []interface{}) {
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

var _ plugin.UninstallPluginActor = new(FakeUninstallPluginActor)
