// This file was generated by counterfeiter
package uifakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/ui"
	. "github.com/cloudfoundry/bosh-cli/ui/table"
)

type FakeUI struct {
	ErrorLinefStub        func(pattern string, args ...interface{})
	errorLinefMutex       sync.RWMutex
	errorLinefArgsForCall []struct {
		pattern string
		args    []interface{}
	}
	PrintLinefStub        func(pattern string, args ...interface{})
	printLinefMutex       sync.RWMutex
	printLinefArgsForCall []struct {
		pattern string
		args    []interface{}
	}
	BeginLinefStub        func(pattern string, args ...interface{})
	beginLinefMutex       sync.RWMutex
	beginLinefArgsForCall []struct {
		pattern string
		args    []interface{}
	}
	EndLinefStub        func(pattern string, args ...interface{})
	endLinefMutex       sync.RWMutex
	endLinefArgsForCall []struct {
		pattern string
		args    []interface{}
	}
	PrintBlockStub        func(string)
	printBlockMutex       sync.RWMutex
	printBlockArgsForCall []struct {
		arg1 string
	}
	PrintErrorBlockStub        func(string)
	printErrorBlockMutex       sync.RWMutex
	printErrorBlockArgsForCall []struct {
		arg1 string
	}
	PrintTableStub        func(Table)
	printTableMutex       sync.RWMutex
	printTableArgsForCall []struct {
		arg1 Table
	}
	AskForTextStub        func(label string) (string, error)
	askForTextMutex       sync.RWMutex
	askForTextArgsForCall []struct {
		label string
	}
	askForTextReturns struct {
		result1 string
		result2 error
	}
	AskForChoiceStub        func(label string, options []string) (int, error)
	askForChoiceMutex       sync.RWMutex
	askForChoiceArgsForCall []struct {
		label   string
		options []string
	}
	askForChoiceReturns struct {
		result1 int
		result2 error
	}
	AskForPasswordStub        func(label string) (string, error)
	askForPasswordMutex       sync.RWMutex
	askForPasswordArgsForCall []struct {
		label string
	}
	askForPasswordReturns struct {
		result1 string
		result2 error
	}
	AskForConfirmationStub        func() error
	askForConfirmationMutex       sync.RWMutex
	askForConfirmationArgsForCall []struct{}
	askForConfirmationReturns     struct {
		result1 error
	}
	IsInteractiveStub        func() bool
	isInteractiveMutex       sync.RWMutex
	isInteractiveArgsForCall []struct{}
	isInteractiveReturns     struct {
		result1 bool
	}
	FlushStub        func()
	flushMutex       sync.RWMutex
	flushArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUI) ErrorLinef(pattern string, args ...interface{}) {
	fake.errorLinefMutex.Lock()
	fake.errorLinefArgsForCall = append(fake.errorLinefArgsForCall, struct {
		pattern string
		args    []interface{}
	}{pattern, args})
	fake.recordInvocation("ErrorLinef", []interface{}{pattern, args})
	fake.errorLinefMutex.Unlock()
	if fake.ErrorLinefStub != nil {
		fake.ErrorLinefStub(pattern, args...)
	}
}

func (fake *FakeUI) ErrorLinefCallCount() int {
	fake.errorLinefMutex.RLock()
	defer fake.errorLinefMutex.RUnlock()
	return len(fake.errorLinefArgsForCall)
}

func (fake *FakeUI) ErrorLinefArgsForCall(i int) (string, []interface{}) {
	fake.errorLinefMutex.RLock()
	defer fake.errorLinefMutex.RUnlock()
	return fake.errorLinefArgsForCall[i].pattern, fake.errorLinefArgsForCall[i].args
}

func (fake *FakeUI) PrintLinef(pattern string, args ...interface{}) {
	fake.printLinefMutex.Lock()
	fake.printLinefArgsForCall = append(fake.printLinefArgsForCall, struct {
		pattern string
		args    []interface{}
	}{pattern, args})
	fake.recordInvocation("PrintLinef", []interface{}{pattern, args})
	fake.printLinefMutex.Unlock()
	if fake.PrintLinefStub != nil {
		fake.PrintLinefStub(pattern, args...)
	}
}

func (fake *FakeUI) PrintLinefCallCount() int {
	fake.printLinefMutex.RLock()
	defer fake.printLinefMutex.RUnlock()
	return len(fake.printLinefArgsForCall)
}

func (fake *FakeUI) PrintLinefArgsForCall(i int) (string, []interface{}) {
	fake.printLinefMutex.RLock()
	defer fake.printLinefMutex.RUnlock()
	return fake.printLinefArgsForCall[i].pattern, fake.printLinefArgsForCall[i].args
}

func (fake *FakeUI) BeginLinef(pattern string, args ...interface{}) {
	fake.beginLinefMutex.Lock()
	fake.beginLinefArgsForCall = append(fake.beginLinefArgsForCall, struct {
		pattern string
		args    []interface{}
	}{pattern, args})
	fake.recordInvocation("BeginLinef", []interface{}{pattern, args})
	fake.beginLinefMutex.Unlock()
	if fake.BeginLinefStub != nil {
		fake.BeginLinefStub(pattern, args...)
	}
}

func (fake *FakeUI) BeginLinefCallCount() int {
	fake.beginLinefMutex.RLock()
	defer fake.beginLinefMutex.RUnlock()
	return len(fake.beginLinefArgsForCall)
}

func (fake *FakeUI) BeginLinefArgsForCall(i int) (string, []interface{}) {
	fake.beginLinefMutex.RLock()
	defer fake.beginLinefMutex.RUnlock()
	return fake.beginLinefArgsForCall[i].pattern, fake.beginLinefArgsForCall[i].args
}

func (fake *FakeUI) EndLinef(pattern string, args ...interface{}) {
	fake.endLinefMutex.Lock()
	fake.endLinefArgsForCall = append(fake.endLinefArgsForCall, struct {
		pattern string
		args    []interface{}
	}{pattern, args})
	fake.recordInvocation("EndLinef", []interface{}{pattern, args})
	fake.endLinefMutex.Unlock()
	if fake.EndLinefStub != nil {
		fake.EndLinefStub(pattern, args...)
	}
}

func (fake *FakeUI) EndLinefCallCount() int {
	fake.endLinefMutex.RLock()
	defer fake.endLinefMutex.RUnlock()
	return len(fake.endLinefArgsForCall)
}

func (fake *FakeUI) EndLinefArgsForCall(i int) (string, []interface{}) {
	fake.endLinefMutex.RLock()
	defer fake.endLinefMutex.RUnlock()
	return fake.endLinefArgsForCall[i].pattern, fake.endLinefArgsForCall[i].args
}

func (fake *FakeUI) PrintBlock(arg1 string) {
	fake.printBlockMutex.Lock()
	fake.printBlockArgsForCall = append(fake.printBlockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("PrintBlock", []interface{}{arg1})
	fake.printBlockMutex.Unlock()
	if fake.PrintBlockStub != nil {
		fake.PrintBlockStub(arg1)
	}
}

func (fake *FakeUI) PrintBlockCallCount() int {
	fake.printBlockMutex.RLock()
	defer fake.printBlockMutex.RUnlock()
	return len(fake.printBlockArgsForCall)
}

func (fake *FakeUI) PrintBlockArgsForCall(i int) string {
	fake.printBlockMutex.RLock()
	defer fake.printBlockMutex.RUnlock()
	return fake.printBlockArgsForCall[i].arg1
}

func (fake *FakeUI) PrintErrorBlock(arg1 string) {
	fake.printErrorBlockMutex.Lock()
	fake.printErrorBlockArgsForCall = append(fake.printErrorBlockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("PrintErrorBlock", []interface{}{arg1})
	fake.printErrorBlockMutex.Unlock()
	if fake.PrintErrorBlockStub != nil {
		fake.PrintErrorBlockStub(arg1)
	}
}

func (fake *FakeUI) PrintErrorBlockCallCount() int {
	fake.printErrorBlockMutex.RLock()
	defer fake.printErrorBlockMutex.RUnlock()
	return len(fake.printErrorBlockArgsForCall)
}

func (fake *FakeUI) PrintErrorBlockArgsForCall(i int) string {
	fake.printErrorBlockMutex.RLock()
	defer fake.printErrorBlockMutex.RUnlock()
	return fake.printErrorBlockArgsForCall[i].arg1
}

func (fake *FakeUI) PrintTable(arg1 Table) {
	fake.printTableMutex.Lock()
	fake.printTableArgsForCall = append(fake.printTableArgsForCall, struct {
		arg1 Table
	}{arg1})
	fake.recordInvocation("PrintTable", []interface{}{arg1})
	fake.printTableMutex.Unlock()
	if fake.PrintTableStub != nil {
		fake.PrintTableStub(arg1)
	}
}

func (fake *FakeUI) PrintTableCallCount() int {
	fake.printTableMutex.RLock()
	defer fake.printTableMutex.RUnlock()
	return len(fake.printTableArgsForCall)
}

func (fake *FakeUI) PrintTableArgsForCall(i int) Table {
	fake.printTableMutex.RLock()
	defer fake.printTableMutex.RUnlock()
	return fake.printTableArgsForCall[i].arg1
}

func (fake *FakeUI) AskForText(label string) (string, error) {
	fake.askForTextMutex.Lock()
	fake.askForTextArgsForCall = append(fake.askForTextArgsForCall, struct {
		label string
	}{label})
	fake.recordInvocation("AskForText", []interface{}{label})
	fake.askForTextMutex.Unlock()
	if fake.AskForTextStub != nil {
		return fake.AskForTextStub(label)
	} else {
		return fake.askForTextReturns.result1, fake.askForTextReturns.result2
	}
}

func (fake *FakeUI) AskForTextCallCount() int {
	fake.askForTextMutex.RLock()
	defer fake.askForTextMutex.RUnlock()
	return len(fake.askForTextArgsForCall)
}

func (fake *FakeUI) AskForTextArgsForCall(i int) string {
	fake.askForTextMutex.RLock()
	defer fake.askForTextMutex.RUnlock()
	return fake.askForTextArgsForCall[i].label
}

func (fake *FakeUI) AskForTextReturns(result1 string, result2 error) {
	fake.AskForTextStub = nil
	fake.askForTextReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUI) AskForChoice(label string, options []string) (int, error) {
	var optionsCopy []string
	if options != nil {
		optionsCopy = make([]string, len(options))
		copy(optionsCopy, options)
	}
	fake.askForChoiceMutex.Lock()
	fake.askForChoiceArgsForCall = append(fake.askForChoiceArgsForCall, struct {
		label   string
		options []string
	}{label, optionsCopy})
	fake.recordInvocation("AskForChoice", []interface{}{label, optionsCopy})
	fake.askForChoiceMutex.Unlock()
	if fake.AskForChoiceStub != nil {
		return fake.AskForChoiceStub(label, options)
	} else {
		return fake.askForChoiceReturns.result1, fake.askForChoiceReturns.result2
	}
}

func (fake *FakeUI) AskForChoiceCallCount() int {
	fake.askForChoiceMutex.RLock()
	defer fake.askForChoiceMutex.RUnlock()
	return len(fake.askForChoiceArgsForCall)
}

func (fake *FakeUI) AskForChoiceArgsForCall(i int) (string, []string) {
	fake.askForChoiceMutex.RLock()
	defer fake.askForChoiceMutex.RUnlock()
	return fake.askForChoiceArgsForCall[i].label, fake.askForChoiceArgsForCall[i].options
}

func (fake *FakeUI) AskForChoiceReturns(result1 int, result2 error) {
	fake.AskForChoiceStub = nil
	fake.askForChoiceReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeUI) AskForPassword(label string) (string, error) {
	fake.askForPasswordMutex.Lock()
	fake.askForPasswordArgsForCall = append(fake.askForPasswordArgsForCall, struct {
		label string
	}{label})
	fake.recordInvocation("AskForPassword", []interface{}{label})
	fake.askForPasswordMutex.Unlock()
	if fake.AskForPasswordStub != nil {
		return fake.AskForPasswordStub(label)
	} else {
		return fake.askForPasswordReturns.result1, fake.askForPasswordReturns.result2
	}
}

func (fake *FakeUI) AskForPasswordCallCount() int {
	fake.askForPasswordMutex.RLock()
	defer fake.askForPasswordMutex.RUnlock()
	return len(fake.askForPasswordArgsForCall)
}

func (fake *FakeUI) AskForPasswordArgsForCall(i int) string {
	fake.askForPasswordMutex.RLock()
	defer fake.askForPasswordMutex.RUnlock()
	return fake.askForPasswordArgsForCall[i].label
}

func (fake *FakeUI) AskForPasswordReturns(result1 string, result2 error) {
	fake.AskForPasswordStub = nil
	fake.askForPasswordReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUI) AskForConfirmation() error {
	fake.askForConfirmationMutex.Lock()
	fake.askForConfirmationArgsForCall = append(fake.askForConfirmationArgsForCall, struct{}{})
	fake.recordInvocation("AskForConfirmation", []interface{}{})
	fake.askForConfirmationMutex.Unlock()
	if fake.AskForConfirmationStub != nil {
		return fake.AskForConfirmationStub()
	} else {
		return fake.askForConfirmationReturns.result1
	}
}

func (fake *FakeUI) AskForConfirmationCallCount() int {
	fake.askForConfirmationMutex.RLock()
	defer fake.askForConfirmationMutex.RUnlock()
	return len(fake.askForConfirmationArgsForCall)
}

func (fake *FakeUI) AskForConfirmationReturns(result1 error) {
	fake.AskForConfirmationStub = nil
	fake.askForConfirmationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUI) IsInteractive() bool {
	fake.isInteractiveMutex.Lock()
	fake.isInteractiveArgsForCall = append(fake.isInteractiveArgsForCall, struct{}{})
	fake.recordInvocation("IsInteractive", []interface{}{})
	fake.isInteractiveMutex.Unlock()
	if fake.IsInteractiveStub != nil {
		return fake.IsInteractiveStub()
	} else {
		return fake.isInteractiveReturns.result1
	}
}

func (fake *FakeUI) IsInteractiveCallCount() int {
	fake.isInteractiveMutex.RLock()
	defer fake.isInteractiveMutex.RUnlock()
	return len(fake.isInteractiveArgsForCall)
}

func (fake *FakeUI) IsInteractiveReturns(result1 bool) {
	fake.IsInteractiveStub = nil
	fake.isInteractiveReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeUI) Flush() {
	fake.flushMutex.Lock()
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct{}{})
	fake.recordInvocation("Flush", []interface{}{})
	fake.flushMutex.Unlock()
	if fake.FlushStub != nil {
		fake.FlushStub()
	}
}

func (fake *FakeUI) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *FakeUI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.errorLinefMutex.RLock()
	defer fake.errorLinefMutex.RUnlock()
	fake.printLinefMutex.RLock()
	defer fake.printLinefMutex.RUnlock()
	fake.beginLinefMutex.RLock()
	defer fake.beginLinefMutex.RUnlock()
	fake.endLinefMutex.RLock()
	defer fake.endLinefMutex.RUnlock()
	fake.printBlockMutex.RLock()
	defer fake.printBlockMutex.RUnlock()
	fake.printErrorBlockMutex.RLock()
	defer fake.printErrorBlockMutex.RUnlock()
	fake.printTableMutex.RLock()
	defer fake.printTableMutex.RUnlock()
	fake.askForTextMutex.RLock()
	defer fake.askForTextMutex.RUnlock()
	fake.askForChoiceMutex.RLock()
	defer fake.askForChoiceMutex.RUnlock()
	fake.askForPasswordMutex.RLock()
	defer fake.askForPasswordMutex.RUnlock()
	fake.askForConfirmationMutex.RLock()
	defer fake.askForConfirmationMutex.RUnlock()
	fake.isInteractiveMutex.RLock()
	defer fake.isInteractiveMutex.RUnlock()
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUI) recordInvocation(key string, args []interface{}) {
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

var _ ui.UI = new(FakeUI)