// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type SigService struct {
	GetSignerStub        func(identity view.Identity) (api.Signer, error)
	getSignerMutex       sync.RWMutex
	getSignerArgsForCall []struct {
		identity view.Identity
	}
	getSignerReturns struct {
		result1 api.Signer
		result2 error
	}
	getSignerReturnsOnCall map[int]struct {
		result1 api.Signer
		result2 error
	}
	GetVerifierStub        func(identity view.Identity) (api.Verifier, error)
	getVerifierMutex       sync.RWMutex
	getVerifierArgsForCall []struct {
		identity view.Identity
	}
	getVerifierReturns struct {
		result1 api.Verifier
		result2 error
	}
	getVerifierReturnsOnCall map[int]struct {
		result1 api.Verifier
		result2 error
	}
	GetSigningIdentityStub        func(identity view.Identity) (api.SigningIdentity, error)
	getSigningIdentityMutex       sync.RWMutex
	getSigningIdentityArgsForCall []struct {
		identity view.Identity
	}
	getSigningIdentityReturns struct {
		result1 api.SigningIdentity
		result2 error
	}
	getSigningIdentityReturnsOnCall map[int]struct {
		result1 api.SigningIdentity
		result2 error
	}
	IdentityTypeStub        func(identity view.Identity) api.IdentityType
	identityTypeMutex       sync.RWMutex
	identityTypeArgsForCall []struct {
		identity view.Identity
	}
	identityTypeReturns struct {
		result1 api.IdentityType
	}
	identityTypeReturnsOnCall map[int]struct {
		result1 api.IdentityType
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SigService) GetSigner(identity view.Identity) (api.Signer, error) {
	fake.getSignerMutex.Lock()
	ret, specificReturn := fake.getSignerReturnsOnCall[len(fake.getSignerArgsForCall)]
	fake.getSignerArgsForCall = append(fake.getSignerArgsForCall, struct {
		identity view.Identity
	}{identity})
	fake.recordInvocation("GetSigner", []interface{}{identity})
	fake.getSignerMutex.Unlock()
	if fake.GetSignerStub != nil {
		return fake.GetSignerStub(identity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSignerReturns.result1, fake.getSignerReturns.result2
}

func (fake *SigService) GetSignerCallCount() int {
	fake.getSignerMutex.RLock()
	defer fake.getSignerMutex.RUnlock()
	return len(fake.getSignerArgsForCall)
}

func (fake *SigService) GetSignerArgsForCall(i int) view.Identity {
	fake.getSignerMutex.RLock()
	defer fake.getSignerMutex.RUnlock()
	return fake.getSignerArgsForCall[i].identity
}

func (fake *SigService) GetSignerReturns(result1 api.Signer, result2 error) {
	fake.GetSignerStub = nil
	fake.getSignerReturns = struct {
		result1 api.Signer
		result2 error
	}{result1, result2}
}

func (fake *SigService) GetSignerReturnsOnCall(i int, result1 api.Signer, result2 error) {
	fake.GetSignerStub = nil
	if fake.getSignerReturnsOnCall == nil {
		fake.getSignerReturnsOnCall = make(map[int]struct {
			result1 api.Signer
			result2 error
		})
	}
	fake.getSignerReturnsOnCall[i] = struct {
		result1 api.Signer
		result2 error
	}{result1, result2}
}

func (fake *SigService) GetVerifier(identity view.Identity) (api.Verifier, error) {
	fake.getVerifierMutex.Lock()
	ret, specificReturn := fake.getVerifierReturnsOnCall[len(fake.getVerifierArgsForCall)]
	fake.getVerifierArgsForCall = append(fake.getVerifierArgsForCall, struct {
		identity view.Identity
	}{identity})
	fake.recordInvocation("GetVerifier", []interface{}{identity})
	fake.getVerifierMutex.Unlock()
	if fake.GetVerifierStub != nil {
		return fake.GetVerifierStub(identity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVerifierReturns.result1, fake.getVerifierReturns.result2
}

func (fake *SigService) GetVerifierCallCount() int {
	fake.getVerifierMutex.RLock()
	defer fake.getVerifierMutex.RUnlock()
	return len(fake.getVerifierArgsForCall)
}

func (fake *SigService) GetVerifierArgsForCall(i int) view.Identity {
	fake.getVerifierMutex.RLock()
	defer fake.getVerifierMutex.RUnlock()
	return fake.getVerifierArgsForCall[i].identity
}

func (fake *SigService) GetVerifierReturns(result1 api.Verifier, result2 error) {
	fake.GetVerifierStub = nil
	fake.getVerifierReturns = struct {
		result1 api.Verifier
		result2 error
	}{result1, result2}
}

func (fake *SigService) GetVerifierReturnsOnCall(i int, result1 api.Verifier, result2 error) {
	fake.GetVerifierStub = nil
	if fake.getVerifierReturnsOnCall == nil {
		fake.getVerifierReturnsOnCall = make(map[int]struct {
			result1 api.Verifier
			result2 error
		})
	}
	fake.getVerifierReturnsOnCall[i] = struct {
		result1 api.Verifier
		result2 error
	}{result1, result2}
}

func (fake *SigService) GetSigningIdentity(identity view.Identity) (api.SigningIdentity, error) {
	fake.getSigningIdentityMutex.Lock()
	ret, specificReturn := fake.getSigningIdentityReturnsOnCall[len(fake.getSigningIdentityArgsForCall)]
	fake.getSigningIdentityArgsForCall = append(fake.getSigningIdentityArgsForCall, struct {
		identity view.Identity
	}{identity})
	fake.recordInvocation("GetSigningIdentity", []interface{}{identity})
	fake.getSigningIdentityMutex.Unlock()
	if fake.GetSigningIdentityStub != nil {
		return fake.GetSigningIdentityStub(identity)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSigningIdentityReturns.result1, fake.getSigningIdentityReturns.result2
}

func (fake *SigService) GetSigningIdentityCallCount() int {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	return len(fake.getSigningIdentityArgsForCall)
}

func (fake *SigService) GetSigningIdentityArgsForCall(i int) view.Identity {
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	return fake.getSigningIdentityArgsForCall[i].identity
}

func (fake *SigService) GetSigningIdentityReturns(result1 api.SigningIdentity, result2 error) {
	fake.GetSigningIdentityStub = nil
	fake.getSigningIdentityReturns = struct {
		result1 api.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *SigService) GetSigningIdentityReturnsOnCall(i int, result1 api.SigningIdentity, result2 error) {
	fake.GetSigningIdentityStub = nil
	if fake.getSigningIdentityReturnsOnCall == nil {
		fake.getSigningIdentityReturnsOnCall = make(map[int]struct {
			result1 api.SigningIdentity
			result2 error
		})
	}
	fake.getSigningIdentityReturnsOnCall[i] = struct {
		result1 api.SigningIdentity
		result2 error
	}{result1, result2}
}

func (fake *SigService) IdentityType(identity view.Identity) api.IdentityType {
	fake.identityTypeMutex.Lock()
	ret, specificReturn := fake.identityTypeReturnsOnCall[len(fake.identityTypeArgsForCall)]
	fake.identityTypeArgsForCall = append(fake.identityTypeArgsForCall, struct {
		identity view.Identity
	}{identity})
	fake.recordInvocation("IdentityType", []interface{}{identity})
	fake.identityTypeMutex.Unlock()
	if fake.IdentityTypeStub != nil {
		return fake.IdentityTypeStub(identity)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.identityTypeReturns.result1
}

func (fake *SigService) IdentityTypeCallCount() int {
	fake.identityTypeMutex.RLock()
	defer fake.identityTypeMutex.RUnlock()
	return len(fake.identityTypeArgsForCall)
}

func (fake *SigService) IdentityTypeArgsForCall(i int) view.Identity {
	fake.identityTypeMutex.RLock()
	defer fake.identityTypeMutex.RUnlock()
	return fake.identityTypeArgsForCall[i].identity
}

func (fake *SigService) IdentityTypeReturns(result1 api.IdentityType) {
	fake.IdentityTypeStub = nil
	fake.identityTypeReturns = struct {
		result1 api.IdentityType
	}{result1}
}

func (fake *SigService) IdentityTypeReturnsOnCall(i int, result1 api.IdentityType) {
	fake.IdentityTypeStub = nil
	if fake.identityTypeReturnsOnCall == nil {
		fake.identityTypeReturnsOnCall = make(map[int]struct {
			result1 api.IdentityType
		})
	}
	fake.identityTypeReturnsOnCall[i] = struct {
		result1 api.IdentityType
	}{result1}
}

func (fake *SigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSignerMutex.RLock()
	defer fake.getSignerMutex.RUnlock()
	fake.getVerifierMutex.RLock()
	defer fake.getVerifierMutex.RUnlock()
	fake.getSigningIdentityMutex.RLock()
	defer fake.getSigningIdentityMutex.RUnlock()
	fake.identityTypeMutex.RLock()
	defer fake.identityTypeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SigService) recordInvocation(key string, args []interface{}) {
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

var _ api.SigService = new(SigService)
