// Code generated by counterfeiter. DO NOT EDIT.
package servicesfakes

import (
	"golang-oauth/services"
	"sync"

	"golang.org/x/oauth2"
)

type FakeCommonInterface struct {
	HandleLoginStub        func(*oauth2.Config, string) (string, error)
	handleLoginMutex       sync.RWMutex
	handleLoginArgsForCall []struct {
		arg1 *oauth2.Config
		arg2 string
	}
	handleLoginReturns struct {
		result1 string
		result2 error
	}
	handleLoginReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommonInterface) HandleLogin(arg1 *oauth2.Config, arg2 string) (string, error) {
	fake.handleLoginMutex.Lock()
	ret, specificReturn := fake.handleLoginReturnsOnCall[len(fake.handleLoginArgsForCall)]
	fake.handleLoginArgsForCall = append(fake.handleLoginArgsForCall, struct {
		arg1 *oauth2.Config
		arg2 string
	}{arg1, arg2})
	stub := fake.HandleLoginStub
	fakeReturns := fake.handleLoginReturns
	fake.recordInvocation("HandleLogin", []interface{}{arg1, arg2})
	fake.handleLoginMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCommonInterface) HandleLoginCallCount() int {
	fake.handleLoginMutex.RLock()
	defer fake.handleLoginMutex.RUnlock()
	return len(fake.handleLoginArgsForCall)
}

func (fake *FakeCommonInterface) HandleLoginCalls(stub func(*oauth2.Config, string) (string, error)) {
	fake.handleLoginMutex.Lock()
	defer fake.handleLoginMutex.Unlock()
	fake.HandleLoginStub = stub
}

func (fake *FakeCommonInterface) HandleLoginArgsForCall(i int) (*oauth2.Config, string) {
	fake.handleLoginMutex.RLock()
	defer fake.handleLoginMutex.RUnlock()
	argsForCall := fake.handleLoginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCommonInterface) HandleLoginReturns(result1 string, result2 error) {
	fake.handleLoginMutex.Lock()
	defer fake.handleLoginMutex.Unlock()
	fake.HandleLoginStub = nil
	fake.handleLoginReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCommonInterface) HandleLoginReturnsOnCall(i int, result1 string, result2 error) {
	fake.handleLoginMutex.Lock()
	defer fake.handleLoginMutex.Unlock()
	fake.HandleLoginStub = nil
	if fake.handleLoginReturnsOnCall == nil {
		fake.handleLoginReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.handleLoginReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCommonInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleLoginMutex.RLock()
	defer fake.handleLoginMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommonInterface) recordInvocation(key string, args []interface{}) {
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

var _ services.CommonInterface = new(FakeCommonInterface)
