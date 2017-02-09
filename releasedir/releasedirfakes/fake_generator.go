// This file was generated by counterfeiter
package releasedirfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeGenerator struct {
	GenerateJobStub        func(string) error
	generateJobMutex       sync.RWMutex
	generateJobArgsForCall []struct {
		arg1 string
	}
	generateJobReturns struct {
		result1 error
	}
	GeneratePackageStub        func(string) error
	generatePackageMutex       sync.RWMutex
	generatePackageArgsForCall []struct {
		arg1 string
	}
	generatePackageReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGenerator) GenerateJob(arg1 string) error {
	fake.generateJobMutex.Lock()
	fake.generateJobArgsForCall = append(fake.generateJobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GenerateJob", []interface{}{arg1})
	fake.generateJobMutex.Unlock()
	if fake.GenerateJobStub != nil {
		return fake.GenerateJobStub(arg1)
	} else {
		return fake.generateJobReturns.result1
	}
}

func (fake *FakeGenerator) GenerateJobCallCount() int {
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	return len(fake.generateJobArgsForCall)
}

func (fake *FakeGenerator) GenerateJobArgsForCall(i int) string {
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	return fake.generateJobArgsForCall[i].arg1
}

func (fake *FakeGenerator) GenerateJobReturns(result1 error) {
	fake.GenerateJobStub = nil
	fake.generateJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) GeneratePackage(arg1 string) error {
	fake.generatePackageMutex.Lock()
	fake.generatePackageArgsForCall = append(fake.generatePackageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GeneratePackage", []interface{}{arg1})
	fake.generatePackageMutex.Unlock()
	if fake.GeneratePackageStub != nil {
		return fake.GeneratePackageStub(arg1)
	} else {
		return fake.generatePackageReturns.result1
	}
}

func (fake *FakeGenerator) GeneratePackageCallCount() int {
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	return len(fake.generatePackageArgsForCall)
}

func (fake *FakeGenerator) GeneratePackageArgsForCall(i int) string {
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	return fake.generatePackageArgsForCall[i].arg1
}

func (fake *FakeGenerator) GeneratePackageReturns(result1 error) {
	fake.GeneratePackageStub = nil
	fake.generatePackageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateJobMutex.RLock()
	defer fake.generateJobMutex.RUnlock()
	fake.generatePackageMutex.RLock()
	defer fake.generatePackageMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeGenerator) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.Generator = new(FakeGenerator)
