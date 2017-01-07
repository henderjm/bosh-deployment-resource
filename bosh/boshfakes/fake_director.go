// This file was generated by counterfeiter
package boshfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-deployment-resource/bosh"
)

type FakeDirector struct {
	DeployStub        func(manifest string) error
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		manifest string
	}
	deployReturns struct {
		result1 error
	}
	DownloadManifestStub        func() ([]byte, error)
	downloadManifestMutex       sync.RWMutex
	downloadManifestArgsForCall []struct{}
	downloadManifestReturns     struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDirector) Deploy(manifest string) error {
	fake.deployMutex.Lock()
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		manifest string
	}{manifest})
	fake.recordInvocation("Deploy", []interface{}{manifest})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(manifest)
	}
	return fake.deployReturns.result1
}

func (fake *FakeDirector) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeDirector) DeployArgsForCall(i int) string {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return fake.deployArgsForCall[i].manifest
}

func (fake *FakeDirector) DeployReturns(result1 error) {
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDirector) DownloadManifest() ([]byte, error) {
	fake.downloadManifestMutex.Lock()
	fake.downloadManifestArgsForCall = append(fake.downloadManifestArgsForCall, struct{}{})
	fake.recordInvocation("DownloadManifest", []interface{}{})
	fake.downloadManifestMutex.Unlock()
	if fake.DownloadManifestStub != nil {
		return fake.DownloadManifestStub()
	}
	return fake.downloadManifestReturns.result1, fake.downloadManifestReturns.result2
}

func (fake *FakeDirector) DownloadManifestCallCount() int {
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	return len(fake.downloadManifestArgsForCall)
}

func (fake *FakeDirector) DownloadManifestReturns(result1 []byte, result2 error) {
	fake.DownloadManifestStub = nil
	fake.downloadManifestReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeDirector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDirector) recordInvocation(key string, args []interface{}) {
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

var _ bosh.Director = new(FakeDirector)
