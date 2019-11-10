// Source: DeclaredInterface.go

// Package AppInterfaces is a generated GoMock package.
package AppInterfaces

import (
	ScopingInterfaces "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	yacc "github.com/bhbosman/CodeGenerators/idlgenerator/yacc"
	gomock "github.com/bhbosman/gomock/gomock"
	io "io"
	reflect "reflect"
)

// MockIProcessInformation is a mock of IProcessInformation interface
type MockIProcessInformation struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIProcessInformationMockRecorder
}

// MockIProcessInformationMockRecorder is the mock recorder for MockIProcessInformation
type MockIProcessInformationMockRecorder struct {
	mock     *MockIProcessInformation
	instance interface{}
}

// NewMockIProcessInformation creates a new mock instance
func NewMockIProcessInformation(ctrl *gomock.Controller) *MockIProcessInformation {
	mock := &MockIProcessInformation{ctrl: ctrl}
	mock.recorder = &MockIProcessInformationMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIProcessInformationInstanceWrapper creates a new mock instance
func NewMockIProcessInformationInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIProcessInformation {
	mock := &MockIProcessInformation{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIProcessInformationMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIProcessInformation) EXPECT() *MockIProcessInformationMockRecorder {
	return m.recorder
}

// GetFileName mocks base method
func (m *MockIProcessInformation) GetFileName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetFileName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFileName indicates an expected call of GetFileName
func (mr *MockIProcessInformationMockRecorder) GetFileName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetFileName",
			MethodType:       reflect.TypeOf((*MockIProcessInformation)(nil).GetFileName),
		})
}

// GetArg mocks base method
func (m *MockIProcessInformation) GetArg() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetArg")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetArg indicates an expected call of GetArg
func (mr *MockIProcessInformationMockRecorder) GetArg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetArg",
			MethodType:       reflect.TypeOf((*MockIProcessInformation)(nil).GetArg),
		})
}

// MockIFileInformation is a mock of IFileInformation interface
type MockIFileInformation struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIFileInformationMockRecorder
}

// MockIFileInformationMockRecorder is the mock recorder for MockIFileInformation
type MockIFileInformationMockRecorder struct {
	mock     *MockIFileInformation
	instance interface{}
}

// NewMockIFileInformation creates a new mock instance
func NewMockIFileInformation(ctrl *gomock.Controller) *MockIFileInformation {
	mock := &MockIFileInformation{ctrl: ctrl}
	mock.recorder = &MockIFileInformationMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIFileInformationInstanceWrapper creates a new mock instance
func NewMockIFileInformationInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIFileInformation {
	mock := &MockIFileInformation{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIFileInformationMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFileInformation) EXPECT() *MockIFileInformationMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIFileInformation) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIFileInformationMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Close",
			MethodType:       reflect.TypeOf((*MockIFileInformation)(nil).Close),
		})
}

// GetFileName mocks base method
func (m *MockIFileInformation) GetFileName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetFileName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFileName indicates an expected call of GetFileName
func (mr *MockIFileInformationMockRecorder) GetFileName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetFileName",
			MethodType:       reflect.TypeOf((*MockIFileInformation)(nil).GetFileName),
		})
}

// GetArg mocks base method
func (m *MockIFileInformation) GetArg() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetArg")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetArg indicates an expected call of GetArg
func (mr *MockIFileInformationMockRecorder) GetArg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetArg",
			MethodType:       reflect.TypeOf((*MockIFileInformation)(nil).GetArg),
		})
}

// GetReader mocks base method
func (m *MockIFileInformation) GetReader() io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetReader")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// GetReader indicates an expected call of GetReader
func (mr *MockIFileInformationMockRecorder) GetReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetReader",
			MethodType:       reflect.TypeOf((*MockIFileInformation)(nil).GetReader),
		})
}

// MockIIoReaders is a mock of IIoReaders interface
type MockIIoReaders struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIIoReadersMockRecorder
}

// MockIIoReadersMockRecorder is the mock recorder for MockIIoReaders
type MockIIoReadersMockRecorder struct {
	mock     *MockIIoReaders
	instance interface{}
}

// NewMockIIoReaders creates a new mock instance
func NewMockIIoReaders(ctrl *gomock.Controller) *MockIIoReaders {
	mock := &MockIIoReaders{ctrl: ctrl}
	mock.recorder = &MockIIoReadersMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIIoReadersInstanceWrapper creates a new mock instance
func NewMockIIoReadersInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIIoReaders {
	mock := &MockIIoReaders{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIIoReadersMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIIoReaders) EXPECT() *MockIIoReadersMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockIIoReaders) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockIIoReadersMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Start",
			MethodType:       reflect.TypeOf((*MockIIoReaders)(nil).Start),
		})
}

// Stop mocks base method
func (m *MockIIoReaders) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockIIoReadersMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Stop",
			MethodType:       reflect.TypeOf((*MockIIoReaders)(nil).Stop),
		})
}

// GetFileInformation mocks base method
func (m *MockIIoReaders) GetFileInformation() []IFileInformation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetFileInformation")
	ret0, _ := ret[0].([]IFileInformation)
	return ret0
}

// GetFileInformation indicates an expected call of GetFileInformation
func (mr *MockIIoReadersMockRecorder) GetFileInformation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetFileInformation",
			MethodType:       reflect.TypeOf((*MockIIoReaders)(nil).GetFileInformation),
		})
}

// MockIDefinitionContextFactory is a mock of IDefinitionContextFactory interface
type MockIDefinitionContextFactory struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIDefinitionContextFactoryMockRecorder
}

// MockIDefinitionContextFactoryMockRecorder is the mock recorder for MockIDefinitionContextFactory
type MockIDefinitionContextFactoryMockRecorder struct {
	mock     *MockIDefinitionContextFactory
	instance interface{}
}

// NewMockIDefinitionContextFactory creates a new mock instance
func NewMockIDefinitionContextFactory(ctrl *gomock.Controller) *MockIDefinitionContextFactory {
	mock := &MockIDefinitionContextFactory{ctrl: ctrl}
	mock.recorder = &MockIDefinitionContextFactoryMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIDefinitionContextFactoryInstanceWrapper creates a new mock instance
func NewMockIDefinitionContextFactoryInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIDefinitionContextFactory {
	mock := &MockIDefinitionContextFactory{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIDefinitionContextFactoryMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDefinitionContextFactory) EXPECT() *MockIDefinitionContextFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIDefinitionContextFactory) Create() yacc.IDefinitionContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Create")
	ret0, _ := ret[0].(yacc.IDefinitionContext)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIDefinitionContextFactoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Create",
			MethodType:       reflect.TypeOf((*MockIDefinitionContextFactory)(nil).Create),
		})
}

// MockIProcessor is a mock of IProcessor interface
type MockIProcessor struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIProcessorMockRecorder
}

// MockIProcessorMockRecorder is the mock recorder for MockIProcessor
type MockIProcessorMockRecorder struct {
	mock     *MockIProcessor
	instance interface{}
}

// NewMockIProcessor creates a new mock instance
func NewMockIProcessor(ctrl *gomock.Controller) *MockIProcessor {
	mock := &MockIProcessor{ctrl: ctrl}
	mock.recorder = &MockIProcessorMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIProcessorInstanceWrapper creates a new mock instance
func NewMockIProcessorInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIProcessor {
	mock := &MockIProcessor{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIProcessorMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIProcessor) EXPECT() *MockIProcessorMockRecorder {
	return m.recorder
}

// Process mocks base method
func (m *MockIProcessor) Process() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Process")
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process
func (mr *MockIProcessorMockRecorder) Process() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Process",
			MethodType:       reflect.TypeOf((*MockIProcessor)(nil).Process),
		})
}

// MockIGenerateCode is a mock of IScopeWalker interface
type MockIGenerateCode struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIGenerateCodeMockRecorder
}

// MockIGenerateCodeMockRecorder is the mock recorder for MockIGenerateCode
type MockIGenerateCodeMockRecorder struct {
	mock     *MockIGenerateCode
	instance interface{}
}

// NewMockIGenerateCode creates a new mock instance
func NewMockIGenerateCode(ctrl *gomock.Controller) *MockIGenerateCode {
	mock := &MockIGenerateCode{ctrl: ctrl}
	mock.recorder = &MockIGenerateCodeMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIGenerateCodeInstanceWrapper creates a new mock instance
func NewMockIGenerateCodeInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIGenerateCode {
	mock := &MockIGenerateCode{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIGenerateCodeMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGenerateCode) EXPECT() *MockIGenerateCodeMockRecorder {
	return m.recorder
}

// InternalGenerate mocks base method
func (m *MockIGenerateCode) Generate(indent int, definition ScopingInterfaces.ITypeSpec, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "InternalGenerate", indent, definition, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// InternalGenerate indicates an expected call of InternalGenerate
func (mr *MockIGenerateCodeMockRecorder) Generate(indent, definition, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "InternalGenerate",
			MethodType:       reflect.TypeOf((*MockIGenerateCode)(nil).Generate),
		}, indent, definition, fileName)
}

// MockIIdlGeneratorFlags is a mock of IIdlGeneratorFlags interface
type MockIIdlGeneratorFlags struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIIdlGeneratorFlagsMockRecorder
}

// MockIIdlGeneratorFlagsMockRecorder is the mock recorder for MockIIdlGeneratorFlags
type MockIIdlGeneratorFlagsMockRecorder struct {
	mock     *MockIIdlGeneratorFlags
	instance interface{}
}

// NewMockIIdlGeneratorFlags creates a new mock instance
func NewMockIIdlGeneratorFlags(ctrl *gomock.Controller) *MockIIdlGeneratorFlags {
	mock := &MockIIdlGeneratorFlags{ctrl: ctrl}
	mock.recorder = &MockIIdlGeneratorFlagsMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIIdlGeneratorFlagsInstanceWrapper creates a new mock instance
func NewMockIIdlGeneratorFlagsInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIIdlGeneratorFlags {
	mock := &MockIIdlGeneratorFlags{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIIdlGeneratorFlagsMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIIdlGeneratorFlags) EXPECT() *MockIIdlGeneratorFlagsMockRecorder {
	return m.recorder
}

// Files mocks base method
func (m *MockIIdlGeneratorFlags) Files() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Files")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Files indicates an expected call of Files
func (mr *MockIIdlGeneratorFlagsMockRecorder) Files() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Files",
			MethodType:       reflect.TypeOf((*MockIIdlGeneratorFlags)(nil).Files),
		})
}

// MockISetIdlGeneratorFlags is a mock of ISetIdlGeneratorFlags interface
type MockISetIdlGeneratorFlags struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockISetIdlGeneratorFlagsMockRecorder
}

// MockISetIdlGeneratorFlagsMockRecorder is the mock recorder for MockISetIdlGeneratorFlags
type MockISetIdlGeneratorFlagsMockRecorder struct {
	mock     *MockISetIdlGeneratorFlags
	instance interface{}
}

// NewMockISetIdlGeneratorFlags creates a new mock instance
func NewMockISetIdlGeneratorFlags(ctrl *gomock.Controller) *MockISetIdlGeneratorFlags {
	mock := &MockISetIdlGeneratorFlags{ctrl: ctrl}
	mock.recorder = &MockISetIdlGeneratorFlagsMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockISetIdlGeneratorFlagsInstanceWrapper creates a new mock instance
func NewMockISetIdlGeneratorFlagsInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockISetIdlGeneratorFlags {
	mock := &MockISetIdlGeneratorFlags{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockISetIdlGeneratorFlagsMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISetIdlGeneratorFlags) EXPECT() *MockISetIdlGeneratorFlagsMockRecorder {
	return m.recorder
}

// SetFiles mocks base method
func (m *MockISetIdlGeneratorFlags) SetFiles(files []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "SetFiles", files)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFiles indicates an expected call of SetFiles
func (mr *MockISetIdlGeneratorFlagsMockRecorder) SetFiles(files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "SetFiles",
			MethodType:       reflect.TypeOf((*MockISetIdlGeneratorFlags)(nil).SetFiles),
		}, files)
}
