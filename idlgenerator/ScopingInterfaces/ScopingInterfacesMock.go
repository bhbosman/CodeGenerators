// Source: ScopingInterfaces.go

// Package ScopingInterfaces is a generated GoMock package.
package ScopingInterfaces

import (
	gomock "github.com/bhbosman/gomock/gomock"
	reflect "reflect"
)

// MockIScopingContext is a mock of IScopingContext interface
type MockIScopingContext struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIScopingContextMockRecorder
}

// MockIScopingContextMockRecorder is the mock recorder for MockIScopingContext
type MockIScopingContextMockRecorder struct {
	mock     *MockIScopingContext
	instance interface{}
}

// NewMockIScopingContext creates a new mock instance
func NewMockIScopingContext(ctrl *gomock.Controller) *MockIScopingContext {
	mock := &MockIScopingContext{ctrl: ctrl}
	mock.recorder = &MockIScopingContextMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIScopingContextInstanceWrapper creates a new mock instance
func NewMockIScopingContextInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIScopingContext {
	mock := &MockIScopingContext{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIScopingContextMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIScopingContext) EXPECT() *MockIScopingContextMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockIScopingContext) Find(name string) (bool, IDeclaredType) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Find", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(IDeclaredType)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockIScopingContextMockRecorder) Find(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Find",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).Find),
		}, name)
}

// Add mocks base method
func (m *MockIScopingContext) Add(name string, structType IDeclaredType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Add", name, structType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockIScopingContextMockRecorder) Add(name, structType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Add",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).Add),
		}, name, structType)
}

// Iterate mocks base method
func (m *MockIScopingContext) Iterate(cb func(string, IDeclaredType) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Iterate", cb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Iterate indicates an expected call of Iterate
func (mr *MockIScopingContextMockRecorder) Iterate(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Iterate",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).Iterate),
		}, cb)
}

// GetFlags mocks base method
func (m *MockIScopingContext) GetFlags() ScopingContextFlags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetFlags")
	ret0, _ := ret[0].(ScopingContextFlags)
	return ret0
}

// GetFlags indicates an expected call of GetFlags
func (mr *MockIScopingContextMockRecorder) GetFlags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetFlags",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).GetFlags),
		})
}

// AddUnresolved mocks base method
func (m *MockIScopingContext) AddUnresolved(name string, information IFileInformation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "AddUnresolved", name, information)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUnresolved indicates an expected call of AddUnresolved
func (mr *MockIScopingContextMockRecorder) AddUnresolved(name, information interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "AddUnresolved",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).AddUnresolved),
		}, name, information)
}

// FindPrimitive mocks base method
func (m *MockIScopingContext) FindTypeSpec(fileInformation IFileInformation, s string) (IDeclaredType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "FindPrimitive", fileInformation, s)
	ret0, _ := ret[0].(IDeclaredType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrimitive indicates an expected call of FindPrimitive
func (mr *MockIScopingContextMockRecorder) FindTypeSpec(fileInformation, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "FindPrimitive",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).FindTypeSpec),
		}, fileInformation, s)
}

// Previous mocks base method
func (m *MockIScopingContext) Previous() IScopingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Previous")
	ret0, _ := ret[0].(IScopingContext)
	return ret0
}

// Previous indicates an expected call of Previous
func (mr *MockIScopingContextMockRecorder) Previous() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Previous",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).Previous),
		})
}

// IterateUnresolved mocks base method
func (m *MockIScopingContext) IterateUnresolved(arg0 func(string, []IFileInformation) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "IterateUnresolved", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateUnresolved indicates an expected call of IterateUnresolved
func (mr *MockIScopingContextMockRecorder) IterateUnresolved(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "IterateUnresolved",
			MethodType:       reflect.TypeOf((*MockIScopingContext)(nil).IterateUnresolved),
		}, arg0)
}

// MockIScopingContextFactory is a mock of IScopingContextFactory interface
type MockIScopingContextFactory struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIScopingContextFactoryMockRecorder
}

// MockIScopingContextFactoryMockRecorder is the mock recorder for MockIScopingContextFactory
type MockIScopingContextFactoryMockRecorder struct {
	mock     *MockIScopingContextFactory
	instance interface{}
}

// NewMockIScopingContextFactory creates a new mock instance
func NewMockIScopingContextFactory(ctrl *gomock.Controller) *MockIScopingContextFactory {
	mock := &MockIScopingContextFactory{ctrl: ctrl}
	mock.recorder = &MockIScopingContextFactoryMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIScopingContextFactoryInstanceWrapper creates a new mock instance
func NewMockIScopingContextFactoryInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIScopingContextFactory {
	mock := &MockIScopingContextFactory{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIScopingContextFactoryMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIScopingContextFactory) EXPECT() *MockIScopingContextFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIScopingContextFactory) Create(prevContext IScopingContext) IScopingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Create", prevContext)
	ret0, _ := ret[0].(IScopingContext)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockIScopingContextFactoryMockRecorder) Create(prevContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Create",
			MethodType:       reflect.TypeOf((*MockIScopingContextFactory)(nil).Create),
		}, prevContext)
}

// CreateEmpty mocks base method
func (m *MockIScopingContextFactory) CreateEmpty(prevContext IScopingContext) IScopingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateEmpty", prevContext)
	ret0, _ := ret[0].(IScopingContext)
	return ret0
}

// CreateEmpty indicates an expected call of CreateEmpty
func (mr *MockIScopingContextFactoryMockRecorder) CreateEmpty(prevContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateEmpty",
			MethodType:       reflect.TypeOf((*MockIScopingContextFactory)(nil).CreateEmpty),
		}, prevContext)
}

// CreateEmptyForExport mocks base method
func (m *MockIScopingContextFactory) CreateEmptyForExport(prevContext IScopingContext) IScopingContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateEmptyForExport", prevContext)
	ret0, _ := ret[0].(IScopingContext)
	return ret0
}

// CreateEmptyForExport indicates an expected call of CreateEmptyForExport
func (mr *MockIScopingContextFactoryMockRecorder) CreateEmptyForExport(prevContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateEmptyForExport",
			MethodType:       reflect.TypeOf((*MockIScopingContextFactory)(nil).CreateEmptyForExport),
		}, prevContext)
}

// MockIScopingService is a mock of IScopingService interface
type MockIScopingService struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIScopingServiceMockRecorder
}

// MockIScopingServiceMockRecorder is the mock recorder for MockIScopingService
type MockIScopingServiceMockRecorder struct {
	mock     *MockIScopingService
	instance interface{}
}

// NewMockIScopingService creates a new mock instance
func NewMockIScopingService(ctrl *gomock.Controller) *MockIScopingService {
	mock := &MockIScopingService{ctrl: ctrl}
	mock.recorder = &MockIScopingServiceMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIScopingServiceInstanceWrapper creates a new mock instance
func NewMockIScopingServiceInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIScopingService {
	mock := &MockIScopingService{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIScopingServiceMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIScopingService) EXPECT() *MockIScopingServiceMockRecorder {
	return m.recorder
}

// Scope mocks base method
func (m *MockIScopingService) Scope(scopingContext IScopingContext, declaredType IDeclaredType) (IScopingContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Scope", scopingContext, declaredType)
	ret0, _ := ret[0].(IScopingContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scope indicates an expected call of Scope
func (mr *MockIScopingServiceMockRecorder) Scope(scopingContext, declaredType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Scope",
			MethodType:       reflect.TypeOf((*MockIScopingService)(nil).Scope),
		}, scopingContext, declaredType)
}

// MockIDefaultTypeService is a mock of IDefaultTypeService interface
type MockIDefaultTypeService struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIDefaultTypeServiceMockRecorder
}

// MockIDefaultTypeServiceMockRecorder is the mock recorder for MockIDefaultTypeService
type MockIDefaultTypeServiceMockRecorder struct {
	mock     *MockIDefaultTypeService
	instance interface{}
}

// NewMockIDefaultTypeService creates a new mock instance
func NewMockIDefaultTypeService(ctrl *gomock.Controller) *MockIDefaultTypeService {
	mock := &MockIDefaultTypeService{ctrl: ctrl}
	mock.recorder = &MockIDefaultTypeServiceMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIDefaultTypeServiceInstanceWrapper creates a new mock instance
func NewMockIDefaultTypeServiceInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIDefaultTypeService {
	mock := &MockIDefaultTypeService{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIDefaultTypeServiceMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDefaultTypeService) EXPECT() *MockIDefaultTypeServiceMockRecorder {
	return m.recorder
}

// Iterate mocks base method
func (m *MockIDefaultTypeService) Iterate(cb func(string, IDeclaredType)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, m.instance, "Iterate", cb)
}

// Iterate indicates an expected call of Iterate
func (mr *MockIDefaultTypeServiceMockRecorder) Iterate(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Iterate",
			MethodType:       reflect.TypeOf((*MockIDefaultTypeService)(nil).Iterate),
		}, cb)
}

// FindOk mocks base method
func (m *MockIDefaultTypeService) FindOk(s string) (IDeclaredType, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "FindOk", s)
	ret0, _ := ret[0].(IDeclaredType)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindOk indicates an expected call of FindOk
func (mr *MockIDefaultTypeServiceMockRecorder) FindOk(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "FindOk",
			MethodType:       reflect.TypeOf((*MockIDefaultTypeService)(nil).FindOk),
		}, s)
}

// Find mocks base method
func (m *MockIDefaultTypeService) Find(s string) IDeclaredType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Find", s)
	ret0, _ := ret[0].(IDeclaredType)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockIDefaultTypeServiceMockRecorder) Find(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Find",
			MethodType:       reflect.TypeOf((*MockIDefaultTypeService)(nil).Find),
		}, s)
}

// MockICheckMemberName is a mock of ICheckMemberName interface
type MockICheckMemberName struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockICheckMemberNameMockRecorder
}

// MockICheckMemberNameMockRecorder is the mock recorder for MockICheckMemberName
type MockICheckMemberNameMockRecorder struct {
	mock     *MockICheckMemberName
	instance interface{}
}

// NewMockICheckMemberName creates a new mock instance
func NewMockICheckMemberName(ctrl *gomock.Controller) *MockICheckMemberName {
	mock := &MockICheckMemberName{ctrl: ctrl}
	mock.recorder = &MockICheckMemberNameMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockICheckMemberNameInstanceWrapper creates a new mock instance
func NewMockICheckMemberNameInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockICheckMemberName {
	mock := &MockICheckMemberName{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockICheckMemberNameMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICheckMemberName) EXPECT() *MockICheckMemberNameMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockICheckMemberName) Add(identifier string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Add", identifier)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockICheckMemberNameMockRecorder) Add(identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Add",
			MethodType:       reflect.TypeOf((*MockICheckMemberName)(nil).Add),
		}, identifier)
}

// Validate mocks base method
func (m *MockICheckMemberName) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockICheckMemberNameMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Validate",
			MethodType:       reflect.TypeOf((*MockICheckMemberName)(nil).Validate),
		})
}

// MockICheckMemberNameFactory is a mock of ICheckMemberNameFactory interface
type MockICheckMemberNameFactory struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockICheckMemberNameFactoryMockRecorder
}

// MockICheckMemberNameFactoryMockRecorder is the mock recorder for MockICheckMemberNameFactory
type MockICheckMemberNameFactoryMockRecorder struct {
	mock     *MockICheckMemberNameFactory
	instance interface{}
}

// NewMockICheckMemberNameFactory creates a new mock instance
func NewMockICheckMemberNameFactory(ctrl *gomock.Controller) *MockICheckMemberNameFactory {
	mock := &MockICheckMemberNameFactory{ctrl: ctrl}
	mock.recorder = &MockICheckMemberNameFactoryMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockICheckMemberNameFactoryInstanceWrapper creates a new mock instance
func NewMockICheckMemberNameFactoryInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockICheckMemberNameFactory {
	mock := &MockICheckMemberNameFactory{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockICheckMemberNameFactoryMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICheckMemberNameFactory) EXPECT() *MockICheckMemberNameFactoryMockRecorder {
	return m.recorder
}
