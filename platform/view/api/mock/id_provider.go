// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type IdentityProvider struct {
	DefaultIdentityStub        func() view.Identity
	defaultIdentityMutex       sync.RWMutex
	defaultIdentityArgsForCall []struct{}
	defaultIdentityReturns     struct {
		result1 view.Identity
	}
	defaultIdentityReturnsOnCall map[int]struct {
		result1 view.Identity
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IdentityProvider) DefaultIdentity() view.Identity {
	fake.defaultIdentityMutex.Lock()
	ret, specificReturn := fake.defaultIdentityReturnsOnCall[len(fake.defaultIdentityArgsForCall)]
	fake.defaultIdentityArgsForCall = append(fake.defaultIdentityArgsForCall, struct{}{})
	fake.recordInvocation("DefaultIdentity", []interface{}{})
	fake.defaultIdentityMutex.Unlock()
	if fake.DefaultIdentityStub != nil {
		return fake.DefaultIdentityStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.defaultIdentityReturns.result1
}

func (fake *IdentityProvider) DefaultIdentityCallCount() int {
	fake.defaultIdentityMutex.RLock()
	defer fake.defaultIdentityMutex.RUnlock()
	return len(fake.defaultIdentityArgsForCall)
}

func (fake *IdentityProvider) DefaultIdentityReturns(result1 view.Identity) {
	fake.DefaultIdentityStub = nil
	fake.defaultIdentityReturns = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) DefaultIdentityReturnsOnCall(i int, result1 view.Identity) {
	fake.DefaultIdentityStub = nil
	if fake.defaultIdentityReturnsOnCall == nil {
		fake.defaultIdentityReturnsOnCall = make(map[int]struct {
			result1 view.Identity
		})
	}
	fake.defaultIdentityReturnsOnCall[i] = struct {
		result1 view.Identity
	}{result1}
}

func (fake *IdentityProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.defaultIdentityMutex.RLock()
	defer fake.defaultIdentityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IdentityProvider) recordInvocation(key string, args []interface{}) {
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

var _ api.IdentityProvider = new(IdentityProvider)
