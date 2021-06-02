// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/core/generic/csp/idemix/handlers"
)

type User struct {
	MakeNymStub        func(handlers.Big, handlers.IssuerPublicKey) (handlers.Ecp, handlers.Big, error)
	makeNymMutex       sync.RWMutex
	makeNymArgsForCall []struct {
		arg1 handlers.Big
		arg2 handlers.IssuerPublicKey
	}
	makeNymReturns struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}
	makeNymReturnsOnCall map[int]struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}
	NewKeyStub        func() (handlers.Big, error)
	newKeyMutex       sync.RWMutex
	newKeyArgsForCall []struct {
	}
	newKeyReturns struct {
		result1 handlers.Big
		result2 error
	}
	newKeyReturnsOnCall map[int]struct {
		result1 handlers.Big
		result2 error
	}
	NewKeyFromBytesStub        func([]byte) (handlers.Big, error)
	newKeyFromBytesMutex       sync.RWMutex
	newKeyFromBytesArgsForCall []struct {
		arg1 []byte
	}
	newKeyFromBytesReturns struct {
		result1 handlers.Big
		result2 error
	}
	newKeyFromBytesReturnsOnCall map[int]struct {
		result1 handlers.Big
		result2 error
	}
	NewPublicNymFromBytesStub        func([]byte) (handlers.Ecp, error)
	newPublicNymFromBytesMutex       sync.RWMutex
	newPublicNymFromBytesArgsForCall []struct {
		arg1 []byte
	}
	newPublicNymFromBytesReturns struct {
		result1 handlers.Ecp
		result2 error
	}
	newPublicNymFromBytesReturnsOnCall map[int]struct {
		result1 handlers.Ecp
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *User) MakeNym(arg1 handlers.Big, arg2 handlers.IssuerPublicKey) (handlers.Ecp, handlers.Big, error) {
	fake.makeNymMutex.Lock()
	ret, specificReturn := fake.makeNymReturnsOnCall[len(fake.makeNymArgsForCall)]
	fake.makeNymArgsForCall = append(fake.makeNymArgsForCall, struct {
		arg1 handlers.Big
		arg2 handlers.IssuerPublicKey
	}{arg1, arg2})
	fake.recordInvocation("MakeNym", []interface{}{arg1, arg2})
	fake.makeNymMutex.Unlock()
	if fake.MakeNymStub != nil {
		return fake.MakeNymStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.makeNymReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *User) MakeNymCallCount() int {
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	return len(fake.makeNymArgsForCall)
}

func (fake *User) MakeNymCalls(stub func(handlers.Big, handlers.IssuerPublicKey) (handlers.Ecp, handlers.Big, error)) {
	fake.makeNymMutex.Lock()
	defer fake.makeNymMutex.Unlock()
	fake.MakeNymStub = stub
}

func (fake *User) MakeNymArgsForCall(i int) (handlers.Big, handlers.IssuerPublicKey) {
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	argsForCall := fake.makeNymArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *User) MakeNymReturns(result1 handlers.Ecp, result2 handlers.Big, result3 error) {
	fake.makeNymMutex.Lock()
	defer fake.makeNymMutex.Unlock()
	fake.MakeNymStub = nil
	fake.makeNymReturns = struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}{result1, result2, result3}
}

func (fake *User) MakeNymReturnsOnCall(i int, result1 handlers.Ecp, result2 handlers.Big, result3 error) {
	fake.makeNymMutex.Lock()
	defer fake.makeNymMutex.Unlock()
	fake.MakeNymStub = nil
	if fake.makeNymReturnsOnCall == nil {
		fake.makeNymReturnsOnCall = make(map[int]struct {
			result1 handlers.Ecp
			result2 handlers.Big
			result3 error
		})
	}
	fake.makeNymReturnsOnCall[i] = struct {
		result1 handlers.Ecp
		result2 handlers.Big
		result3 error
	}{result1, result2, result3}
}

func (fake *User) NewKey() (handlers.Big, error) {
	fake.newKeyMutex.Lock()
	ret, specificReturn := fake.newKeyReturnsOnCall[len(fake.newKeyArgsForCall)]
	fake.newKeyArgsForCall = append(fake.newKeyArgsForCall, struct {
	}{})
	fake.recordInvocation("NewKey", []interface{}{})
	fake.newKeyMutex.Unlock()
	if fake.NewKeyStub != nil {
		return fake.NewKeyStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *User) NewKeyCallCount() int {
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	return len(fake.newKeyArgsForCall)
}

func (fake *User) NewKeyCalls(stub func() (handlers.Big, error)) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = stub
}

func (fake *User) NewKeyReturns(result1 handlers.Big, result2 error) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = nil
	fake.newKeyReturns = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyReturnsOnCall(i int, result1 handlers.Big, result2 error) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = nil
	if fake.newKeyReturnsOnCall == nil {
		fake.newKeyReturnsOnCall = make(map[int]struct {
			result1 handlers.Big
			result2 error
		})
	}
	fake.newKeyReturnsOnCall[i] = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyFromBytes(arg1 []byte) (handlers.Big, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.newKeyFromBytesMutex.Lock()
	ret, specificReturn := fake.newKeyFromBytesReturnsOnCall[len(fake.newKeyFromBytesArgsForCall)]
	fake.newKeyFromBytesArgsForCall = append(fake.newKeyFromBytesArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("NewKeyFromBytes", []interface{}{arg1Copy})
	fake.newKeyFromBytesMutex.Unlock()
	if fake.NewKeyFromBytesStub != nil {
		return fake.NewKeyFromBytesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newKeyFromBytesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *User) NewKeyFromBytesCallCount() int {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	return len(fake.newKeyFromBytesArgsForCall)
}

func (fake *User) NewKeyFromBytesCalls(stub func([]byte) (handlers.Big, error)) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = stub
}

func (fake *User) NewKeyFromBytesArgsForCall(i int) []byte {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	argsForCall := fake.newKeyFromBytesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *User) NewKeyFromBytesReturns(result1 handlers.Big, result2 error) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = nil
	fake.newKeyFromBytesReturns = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewKeyFromBytesReturnsOnCall(i int, result1 handlers.Big, result2 error) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = nil
	if fake.newKeyFromBytesReturnsOnCall == nil {
		fake.newKeyFromBytesReturnsOnCall = make(map[int]struct {
			result1 handlers.Big
			result2 error
		})
	}
	fake.newKeyFromBytesReturnsOnCall[i] = struct {
		result1 handlers.Big
		result2 error
	}{result1, result2}
}

func (fake *User) NewPublicNymFromBytes(arg1 []byte) (handlers.Ecp, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.newPublicNymFromBytesMutex.Lock()
	ret, specificReturn := fake.newPublicNymFromBytesReturnsOnCall[len(fake.newPublicNymFromBytesArgsForCall)]
	fake.newPublicNymFromBytesArgsForCall = append(fake.newPublicNymFromBytesArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("NewPublicNymFromBytes", []interface{}{arg1Copy})
	fake.newPublicNymFromBytesMutex.Unlock()
	if fake.NewPublicNymFromBytesStub != nil {
		return fake.NewPublicNymFromBytesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newPublicNymFromBytesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *User) NewPublicNymFromBytesCallCount() int {
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	return len(fake.newPublicNymFromBytesArgsForCall)
}

func (fake *User) NewPublicNymFromBytesCalls(stub func([]byte) (handlers.Ecp, error)) {
	fake.newPublicNymFromBytesMutex.Lock()
	defer fake.newPublicNymFromBytesMutex.Unlock()
	fake.NewPublicNymFromBytesStub = stub
}

func (fake *User) NewPublicNymFromBytesArgsForCall(i int) []byte {
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	argsForCall := fake.newPublicNymFromBytesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *User) NewPublicNymFromBytesReturns(result1 handlers.Ecp, result2 error) {
	fake.newPublicNymFromBytesMutex.Lock()
	defer fake.newPublicNymFromBytesMutex.Unlock()
	fake.NewPublicNymFromBytesStub = nil
	fake.newPublicNymFromBytesReturns = struct {
		result1 handlers.Ecp
		result2 error
	}{result1, result2}
}

func (fake *User) NewPublicNymFromBytesReturnsOnCall(i int, result1 handlers.Ecp, result2 error) {
	fake.newPublicNymFromBytesMutex.Lock()
	defer fake.newPublicNymFromBytesMutex.Unlock()
	fake.NewPublicNymFromBytesStub = nil
	if fake.newPublicNymFromBytesReturnsOnCall == nil {
		fake.newPublicNymFromBytesReturnsOnCall = make(map[int]struct {
			result1 handlers.Ecp
			result2 error
		})
	}
	fake.newPublicNymFromBytesReturnsOnCall[i] = struct {
		result1 handlers.Ecp
		result2 error
	}{result1, result2}
}

func (fake *User) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeNymMutex.RLock()
	defer fake.makeNymMutex.RUnlock()
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	fake.newPublicNymFromBytesMutex.RLock()
	defer fake.newPublicNymFromBytesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *User) recordInvocation(key string, args []interface{}) {
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

var _ handlers.User = new(User)
