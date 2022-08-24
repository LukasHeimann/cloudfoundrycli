// Code generated by counterfeiter. DO NOT EDIT.
package wrapperfakes

import (
	"sync"

	"github.com/LukasHeimann/cloudfoundrycli/v8/api/cloudcontroller/wrapper"
)

type FakeTokenCache struct {
	AccessTokenStub        func() string
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct {
	}
	accessTokenReturns struct {
		result1 string
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
	}
	RefreshTokenStub        func() string
	refreshTokenMutex       sync.RWMutex
	refreshTokenArgsForCall []struct {
	}
	refreshTokenReturns struct {
		result1 string
	}
	refreshTokenReturnsOnCall map[int]struct {
		result1 string
	}
	SetAccessTokenStub        func(string)
	setAccessTokenMutex       sync.RWMutex
	setAccessTokenArgsForCall []struct {
		arg1 string
	}
	SetRefreshTokenStub        func(string)
	setRefreshTokenMutex       sync.RWMutex
	setRefreshTokenArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenCache) AccessToken() string {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.accessTokenReturns
	return fakeReturns.result1
}

func (fake *FakeTokenCache) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeTokenCache) AccessTokenCalls(stub func() string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = stub
}

func (fake *FakeTokenCache) AccessTokenReturns(result1 string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) AccessTokenReturnsOnCall(i int, result1 string) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) RefreshToken() string {
	fake.refreshTokenMutex.Lock()
	ret, specificReturn := fake.refreshTokenReturnsOnCall[len(fake.refreshTokenArgsForCall)]
	fake.refreshTokenArgsForCall = append(fake.refreshTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("RefreshToken", []interface{}{})
	fake.refreshTokenMutex.Unlock()
	if fake.RefreshTokenStub != nil {
		return fake.RefreshTokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.refreshTokenReturns
	return fakeReturns.result1
}

func (fake *FakeTokenCache) RefreshTokenCallCount() int {
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	return len(fake.refreshTokenArgsForCall)
}

func (fake *FakeTokenCache) RefreshTokenCalls(stub func() string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = stub
}

func (fake *FakeTokenCache) RefreshTokenReturns(result1 string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = nil
	fake.refreshTokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) RefreshTokenReturnsOnCall(i int, result1 string) {
	fake.refreshTokenMutex.Lock()
	defer fake.refreshTokenMutex.Unlock()
	fake.RefreshTokenStub = nil
	if fake.refreshTokenReturnsOnCall == nil {
		fake.refreshTokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.refreshTokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeTokenCache) SetAccessToken(arg1 string) {
	fake.setAccessTokenMutex.Lock()
	fake.setAccessTokenArgsForCall = append(fake.setAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetAccessToken", []interface{}{arg1})
	fake.setAccessTokenMutex.Unlock()
	if fake.SetAccessTokenStub != nil {
		fake.SetAccessTokenStub(arg1)
	}
}

func (fake *FakeTokenCache) SetAccessTokenCallCount() int {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	return len(fake.setAccessTokenArgsForCall)
}

func (fake *FakeTokenCache) SetAccessTokenCalls(stub func(string)) {
	fake.setAccessTokenMutex.Lock()
	defer fake.setAccessTokenMutex.Unlock()
	fake.SetAccessTokenStub = stub
}

func (fake *FakeTokenCache) SetAccessTokenArgsForCall(i int) string {
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	argsForCall := fake.setAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTokenCache) SetRefreshToken(arg1 string) {
	fake.setRefreshTokenMutex.Lock()
	fake.setRefreshTokenArgsForCall = append(fake.setRefreshTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetRefreshToken", []interface{}{arg1})
	fake.setRefreshTokenMutex.Unlock()
	if fake.SetRefreshTokenStub != nil {
		fake.SetRefreshTokenStub(arg1)
	}
}

func (fake *FakeTokenCache) SetRefreshTokenCallCount() int {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	return len(fake.setRefreshTokenArgsForCall)
}

func (fake *FakeTokenCache) SetRefreshTokenCalls(stub func(string)) {
	fake.setRefreshTokenMutex.Lock()
	defer fake.setRefreshTokenMutex.Unlock()
	fake.SetRefreshTokenStub = stub
}

func (fake *FakeTokenCache) SetRefreshTokenArgsForCall(i int) string {
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	argsForCall := fake.setRefreshTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTokenCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	fake.refreshTokenMutex.RLock()
	defer fake.refreshTokenMutex.RUnlock()
	fake.setAccessTokenMutex.RLock()
	defer fake.setAccessTokenMutex.RUnlock()
	fake.setRefreshTokenMutex.RLock()
	defer fake.setRefreshTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenCache) recordInvocation(key string, args []interface{}) {
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

var _ wrapper.TokenCache = new(FakeTokenCache)
