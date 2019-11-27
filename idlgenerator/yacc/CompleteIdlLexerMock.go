// Source: CompleteIdlLexer.go

// Package yacc is a generated GoMock package.
package yacc

import (
	ScopingInterfaces "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	gomock "github.com/bhbosman/gomock/gomock"
	reflect "reflect"
)

// MockCompleteIdlLexer is a mock of CompleteIdlLexer interface
type MockCompleteIdlLexer struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockCompleteIdlLexerMockRecorder
}

// MockCompleteIdlLexerMockRecorder is the mock recorder for MockCompleteIdlLexer
type MockCompleteIdlLexerMockRecorder struct {
	mock     *MockCompleteIdlLexer
	instance interface{}
}

// NewMockCompleteIdlLexer creates a new mock instance
func NewMockCompleteIdlLexer(ctrl *gomock.Controller) *MockCompleteIdlLexer {
	mock := &MockCompleteIdlLexer{ctrl: ctrl}
	mock.recorder = &MockCompleteIdlLexerMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockCompleteIdlLexerInstanceWrapper creates a new mock instance
func NewMockCompleteIdlLexerInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockCompleteIdlLexer {
	mock := &MockCompleteIdlLexer{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockCompleteIdlLexerMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompleteIdlLexer) EXPECT() *MockCompleteIdlLexerMockRecorder {
	return m.recorder
}

// Lex mocks base method
func (m *MockCompleteIdlLexer) Lex(lval *CompleteIdlSymType) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "Lex", lval)
	ret0, _ := ret[0].(int)
	return ret0
}

// Lex indicates an expected call of Lex
func (mr *MockCompleteIdlLexerMockRecorder) Lex(lval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Lex",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).Lex),
		}, lval)
}

// Error mocks base method
func (m *MockCompleteIdlLexer) Error(s string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, m.instance, "Error", s)
}

// Error indicates an expected call of Error
func (mr *MockCompleteIdlLexerMockRecorder) Error(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "Error",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).Error),
		}, s)
}

// FindPrimitive mocks base method
func (m *MockCompleteIdlLexer) FindPrimitive(fileInformation ScopingInterfaces.IFileInformation, s string) (ScopingInterfaces.IBaseDeclaredType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "FindPrimitive", fileInformation, s)
	ret0, _ := ret[0].(ScopingInterfaces.IBaseDeclaredType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPrimitive indicates an expected call of FindPrimitive
func (mr *MockCompleteIdlLexerMockRecorder) FindPrimitive(fileInformation, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "FindPrimitive",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).FindPrimitive),
		}, fileInformation, s)
}

// InfoAt mocks base method
func (m *MockCompleteIdlLexer) InfoAt(info string, params ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{info}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, m.instance, "InfoAt", varargs...)
}

// InfoAt indicates an expected call of InfoAt
func (mr *MockCompleteIdlLexerMockRecorder) InfoAt(info interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{info}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "InfoAt",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).InfoAt),
		}, varargs...)
}

// LastError mocks base method
func (m *MockCompleteIdlLexer) LastError() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "LastError")
	ret0, _ := ret[0].(string)
	return ret0
}

// LastError indicates an expected call of LastError
func (mr *MockCompleteIdlLexerMockRecorder) LastError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "LastError",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).LastError),
		})
}

// MultiExpr mocks base method
func (m *MockCompleteIdlLexer) MultiExpr(a int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "MultiExpr", a)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiExpr indicates an expected call of MultiExpr
func (mr *MockCompleteIdlLexerMockRecorder) MultiExpr(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "MultiExpr",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).MultiExpr),
		}, a)
}

// AddExpr mocks base method
func (m *MockCompleteIdlLexer) AddExpr(a int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "AddExpr", a)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddExpr indicates an expected call of AddExpr
func (mr *MockCompleteIdlLexerMockRecorder) AddExpr(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "AddExpr",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).AddExpr),
		}, a)
}

// AddOperator mocks base method
func (m *MockCompleteIdlLexer) AddOperator(a, b int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "AddOperator", a, b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOperator indicates an expected call of AddOperator
func (mr *MockCompleteIdlLexerMockRecorder) AddOperator(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "AddOperator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).AddOperator),
		}, a, b)
}

// MinusOperator mocks base method
func (m *MockCompleteIdlLexer) MinusOperator(a, b int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "MinusOperator", a, b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinusOperator indicates an expected call of MinusOperator
func (mr *MockCompleteIdlLexerMockRecorder) MinusOperator(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "MinusOperator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).MinusOperator),
		}, a, b)
}

// DivideOperator mocks base method
func (m *MockCompleteIdlLexer) DivideOperator(a, b int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "DivideOperator", a, b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DivideOperator indicates an expected call of DivideOperator
func (mr *MockCompleteIdlLexerMockRecorder) DivideOperator(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "DivideOperator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).DivideOperator),
		}, a, b)
}

// MultiplyOperator mocks base method
func (m *MockCompleteIdlLexer) MultiplyOperator(a, b int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "MultiplyOperator", a, b)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiplyOperator indicates an expected call of MultiplyOperator
func (mr *MockCompleteIdlLexerMockRecorder) MultiplyOperator(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "MultiplyOperator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).MultiplyOperator),
		}, a, b)
}

// NewStructType mocks base method
func (m *MockCompleteIdlLexer) NewStructType(id ScopingInterfaces.IFileInformation, identifier string, member ScopingInterfaces.IStructMember, forward bool) (ScopingInterfaces.IStructType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "NewStructType", id, identifier, member, forward)
	ret0, _ := ret[0].(ScopingInterfaces.IStructType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStructType indicates an expected call of NewStructType
func (mr *MockCompleteIdlLexerMockRecorder) NewStructType(id, identifier, member, forward interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "NewStructType",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).NewStructType),
		}, id, identifier, member, forward)
}

// TransformString mocks base method
func (m *MockCompleteIdlLexer) TransformString(arg0 string) (ScopingInterfaces.IPrimaryExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "TransformString", arg0)
	ret0, _ := ret[0].(ScopingInterfaces.IPrimaryExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformString indicates an expected call of TransformString
func (mr *MockCompleteIdlLexerMockRecorder) TransformString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "TransformString",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).TransformString),
		}, arg0)
}

// TransformInteger mocks base method
func (m *MockCompleteIdlLexer) TransformInteger(arg0 int) (ScopingInterfaces.IPrimaryExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "TransformInteger", arg0)
	ret0, _ := ret[0].(ScopingInterfaces.IPrimaryExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformInteger indicates an expected call of TransformInteger
func (mr *MockCompleteIdlLexerMockRecorder) TransformInteger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "TransformInteger",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).TransformInteger),
		}, arg0)
}

// TransformValue mocks base method
func (m *MockCompleteIdlLexer) TransformValue(arg0 interface{}, arg1 ScopingInterfaces.IPrimaryExpressionType) (ScopingInterfaces.IPrimaryExpression, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "TransformValue", arg0, arg1)
	ret0, _ := ret[0].(ScopingInterfaces.IPrimaryExpression)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransformValue indicates an expected call of TransformValue
func (mr *MockCompleteIdlLexerMockRecorder) TransformValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "TransformValue",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).TransformValue),
		}, arg0, arg1)
}

// CreateTypePrefixDcl mocks base method
func (m *MockCompleteIdlLexer) CreateTypePrefixDcl(fileInformation ScopingInterfaces.IFileInformation, scopedName, stringLiteral string) (ScopingInterfaces.ITypeSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateTypePrefixDcl", fileInformation, scopedName, stringLiteral)
	ret0, _ := ret[0].(ScopingInterfaces.ITypeSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTypePrefixDcl indicates an expected call of CreateTypePrefixDcl
func (mr *MockCompleteIdlLexerMockRecorder) CreateTypePrefixDcl(fileInformation, scopedName, stringLiteral interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateTypePrefixDcl",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).CreateTypePrefixDcl),
		}, fileInformation, scopedName, stringLiteral)
}

// CreateInterfaceDcl mocks base method
func (m *MockCompleteIdlLexer) CreateInterfaceDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, forward, abstract, local bool, body ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateInterfaceDcl", fileInformation, identifier, forward, abstract, local, body)
	ret0, _ := ret[0].(ScopingInterfaces.ITypeSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInterfaceDcl indicates an expected call of CreateInterfaceDcl
func (mr *MockCompleteIdlLexerMockRecorder) CreateInterfaceDcl(fileInformation, identifier, forward, abstract, local, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateInterfaceDcl",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).CreateInterfaceDcl),
		}, fileInformation, identifier, forward, abstract, local, body)
}

// CreateModuleDcl mocks base method
func (m *MockCompleteIdlLexer) CreateModuleDcl(fileInformation ScopingInterfaces.IFileInformation, moduleName string, typeSpec ScopingInterfaces.ITypeSpec) (ScopingInterfaces.IIdlModuleDcl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateModuleDcl", fileInformation, moduleName, typeSpec)
	ret0, _ := ret[0].(ScopingInterfaces.IIdlModuleDcl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateModuleDcl indicates an expected call of CreateModuleDcl
func (mr *MockCompleteIdlLexerMockRecorder) CreateModuleDcl(fileInformation, moduleName, typeSpec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateModuleDcl",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).CreateModuleDcl),
		}, fileInformation, moduleName, typeSpec)
}

// AssignSpec mocks base method
func (m *MockCompleteIdlLexer) AssignSpec(definition ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "AssignSpec", definition)
	ret0, _ := ret[0].(ScopingInterfaces.ITypeSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignSpec indicates an expected call of AssignSpec
func (mr *MockCompleteIdlLexerMockRecorder) AssignSpec(definition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "AssignSpec",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).AssignSpec),
		}, definition)
}

// GetSpec mocks base method
func (m *MockCompleteIdlLexer) GetSpec() (ScopingInterfaces.ITypeSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "GetSpec")
	ret0, _ := ret[0].(ScopingInterfaces.ITypeSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpec indicates an expected call of GetSpec
func (mr *MockCompleteIdlLexerMockRecorder) GetSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "GetSpec",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).GetSpec),
		})
}

// NewDeclarator mocks base method
func (m *MockCompleteIdlLexer) NewDeclarator(fileInformation ScopingInterfaces.IFileInformation, identifier string) (ScopingInterfaces.IDeclarator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "NewDeclarator", fileInformation, identifier)
	ret0, _ := ret[0].(ScopingInterfaces.IDeclarator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDeclarator indicates an expected call of NewDeclarator
func (mr *MockCompleteIdlLexerMockRecorder) NewDeclarator(fileInformation, identifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "NewDeclarator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).NewDeclarator),
		}, fileInformation, identifier)
}

// CreateInterfaceKind mocks base method
func (m *MockCompleteIdlLexer) CreateInterfaceKind(fileInformation ScopingInterfaces.IFileInformation, local, abstract bool) (ScopingInterfaces.IInterfaceKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "CreateInterfaceKind", fileInformation, local, abstract)
	ret0, _ := ret[0].(ScopingInterfaces.IInterfaceKind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInterfaceKind indicates an expected call of CreateInterfaceKind
func (mr *MockCompleteIdlLexerMockRecorder) CreateInterfaceKind(fileInformation, local, abstract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "CreateInterfaceKind",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).CreateInterfaceKind),
		}, fileInformation, local, abstract)
}

// NewMember mocks base method
func (m *MockCompleteIdlLexer) NewMember(typeSpec ScopingInterfaces.IBaseDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.IStructMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "NewMember", typeSpec, declarator)
	ret0, _ := ret[0].(ScopingInterfaces.IStructMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMember indicates an expected call of NewMember
func (mr *MockCompleteIdlLexerMockRecorder) NewMember(typeSpec, declarator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "NewMember",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).NewMember),
		}, typeSpec, declarator)
}

// NewIdlConstDcl mocks base method
func (m *MockCompleteIdlLexer) NewIdlConstDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, value int) (ScopingInterfaces.IIdlConstDcl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "NewIdlConstDcl", fileInformation, identifier, value)
	ret0, _ := ret[0].(ScopingInterfaces.IIdlConstDcl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewIdlConstDcl indicates an expected call of NewIdlConstDcl
func (mr *MockCompleteIdlLexerMockRecorder) NewIdlConstDcl(fileInformation, identifier, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "NewIdlConstDcl",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).NewIdlConstDcl),
		}, fileInformation, identifier, value)
}

// NewTypeDeclarator mocks base method
func (m *MockCompleteIdlLexer) NewTypeDeclarator(simpleTypeSpec ScopingInterfaces.IBaseDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.ITypeDeclarator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "NewTypeDeclarator", simpleTypeSpec, declarator)
	ret0, _ := ret[0].(ScopingInterfaces.ITypeDeclarator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTypeDeclarator indicates an expected call of NewTypeDeclarator
func (mr *MockCompleteIdlLexerMockRecorder) NewTypeDeclarator(simpleTypeSpec, declarator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "NewTypeDeclarator",
			MethodType:       reflect.TypeOf((*MockCompleteIdlLexer)(nil).NewTypeDeclarator),
		}, simpleTypeSpec, declarator)
}

// MockIDefinitionContext is a mock of IDefinitionContext interface
type MockIDefinitionContext struct {
	ctrl     *gomock.Controller
	instance interface{}
	recorder *MockIDefinitionContextMockRecorder
}

// MockIDefinitionContextMockRecorder is the mock recorder for MockIDefinitionContext
type MockIDefinitionContextMockRecorder struct {
	mock     *MockIDefinitionContext
	instance interface{}
}

// NewMockIDefinitionContext creates a new mock instance
func NewMockIDefinitionContext(ctrl *gomock.Controller) *MockIDefinitionContext {
	mock := &MockIDefinitionContext{ctrl: ctrl}
	mock.recorder = &MockIDefinitionContextMockRecorder{
		mock:     mock,
		instance: nil,
	}
	return mock
}

// NewMockIDefinitionContextInstanceWrapper creates a new mock instance
func NewMockIDefinitionContextInstanceWrapper(ctrl *gomock.Controller, instance interface{}) *MockIDefinitionContext {
	mock := &MockIDefinitionContext{
		ctrl:     ctrl,
		instance: instance,
	}
	mock.recorder = &MockIDefinitionContextMockRecorder{
		mock:     mock,
		instance: instance,
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDefinitionContext) EXPECT() *MockIDefinitionContextMockRecorder {
	return m.recorder
}

// ParseExpression mocks base method
func (m *MockIDefinitionContext) ParseExpression(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, m.instance, "ParseExpression", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ParseExpression indicates an expected call of ParseExpression
func (mr *MockIDefinitionContextMockRecorder) ParseExpression(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "ParseExpression",
			MethodType:       reflect.TypeOf((*MockIDefinitionContext)(nil).ParseExpression),
		}, arg0)
}

// ParseDefinition mocks base method
func (m *MockIDefinitionContext) ParseDefinition(arg0 string, arg1 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, m.instance, "ParseDefinition", arg0, arg1)
}

// ParseDefinition indicates an expected call of ParseDefinition
func (mr *MockIDefinitionContextMockRecorder) ParseDefinition(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "ParseDefinition",
			MethodType:       reflect.TypeOf((*MockIDefinitionContext)(nil).ParseDefinition),
		}, arg0, arg1)
}

// ParsePragma mocks base method
func (m *MockIDefinitionContext) ParsePragma(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, m.instance, "ParsePragma", arg0)
}

// ParsePragma indicates an expected call of ParsePragma
func (mr *MockIDefinitionContextMockRecorder) ParsePragma(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		gomock.ReceiverInfo{
			InstanceReceiver: mr.instance,
			MockReceiver:     mr.mock,
			Method:           "ParsePragma",
			MethodType:       reflect.TypeOf((*MockIDefinitionContext)(nil).ParsePragma),
		}, arg0)
}
