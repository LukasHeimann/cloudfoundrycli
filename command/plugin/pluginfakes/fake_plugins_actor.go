// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/actor/pluginaction"
	"github.com/LukasHeimann/cloudfoundrycli/v8/command/plugin"
)

type FakePluginsActor struct {
	GetOutdatedPluginsStub        func() ([]pluginaction.OutdatedPlugin, error)
	getOutdatedPluginsMutex       sync.RWMutex
	getOutdatedPluginsArgsForCall []struct {
	}
	getOutdatedPluginsReturns struct {
		result1 []pluginaction.OutdatedPlugin
		result2 error
	}
	getOutdatedPluginsReturnsOnCall map[int]struct {
		result1 []pluginaction.OutdatedPlugin
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePluginsActor) GetOutdatedPlugins() ([]pluginaction.OutdatedPlugin, error) {
	fake.getOutdatedPluginsMutex.Lock()
	ret, specificReturn := fake.getOutdatedPluginsReturnsOnCall[len(fake.getOutdatedPluginsArgsForCall)]
	fake.getOutdatedPluginsArgsForCall = append(fake.getOutdatedPluginsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetOutdatedPlugins", []interface{}{})
	fake.getOutdatedPluginsMutex.Unlock()
	if fake.GetOutdatedPluginsStub != nil {
		return fake.GetOutdatedPluginsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getOutdatedPluginsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePluginsActor) GetOutdatedPluginsCallCount() int {
	fake.getOutdatedPluginsMutex.RLock()
	defer fake.getOutdatedPluginsMutex.RUnlock()
	return len(fake.getOutdatedPluginsArgsForCall)
}

func (fake *FakePluginsActor) GetOutdatedPluginsCalls(stub func() ([]pluginaction.OutdatedPlugin, error)) {
	fake.getOutdatedPluginsMutex.Lock()
	defer fake.getOutdatedPluginsMutex.Unlock()
	fake.GetOutdatedPluginsStub = stub
}

func (fake *FakePluginsActor) GetOutdatedPluginsReturns(result1 []pluginaction.OutdatedPlugin, result2 error) {
	fake.getOutdatedPluginsMutex.Lock()
	defer fake.getOutdatedPluginsMutex.Unlock()
	fake.GetOutdatedPluginsStub = nil
	fake.getOutdatedPluginsReturns = struct {
		result1 []pluginaction.OutdatedPlugin
		result2 error
	}{result1, result2}
}

func (fake *FakePluginsActor) GetOutdatedPluginsReturnsOnCall(i int, result1 []pluginaction.OutdatedPlugin, result2 error) {
	fake.getOutdatedPluginsMutex.Lock()
	defer fake.getOutdatedPluginsMutex.Unlock()
	fake.GetOutdatedPluginsStub = nil
	if fake.getOutdatedPluginsReturnsOnCall == nil {
		fake.getOutdatedPluginsReturnsOnCall = make(map[int]struct {
			result1 []pluginaction.OutdatedPlugin
			result2 error
		})
	}
	fake.getOutdatedPluginsReturnsOnCall[i] = struct {
		result1 []pluginaction.OutdatedPlugin
		result2 error
	}{result1, result2}
}

func (fake *FakePluginsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOutdatedPluginsMutex.RLock()
	defer fake.getOutdatedPluginsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePluginsActor) recordInvocation(key string, args []interface{}) {
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

var _ plugin.PluginsActor = new(FakePluginsActor)
