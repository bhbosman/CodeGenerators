package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type OperationDeclarations struct {
	TypeSpecBase
	declarator    ScopingInterfaces.IDeclaredType
	params        ScopingInterfaces.IParameterDeclarations
	exceptionList interface{}
}

func NewOperationDeclarations(
	fileInformation ScopingInterfaces.IFileInformation,
	operationName string,
	declarator ScopingInterfaces.IDeclaredType,
	params ScopingInterfaces.IParameterDeclarations,
	exceptionList interface{}) *OperationDeclarations {

	return &OperationDeclarations{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			operationName,
			ScopingInterfaces.Op_dclIdlType,
			false,
			false,
			false,
			false),
		declarator:    declarator,
		params:        params,
		exceptionList: exceptionList,
	}
}

func (self *OperationDeclarations) GetOperationDeclaratorType() ScopingInterfaces.IDeclaredType {
	return self.declarator
}

func (self *OperationDeclarations) GetOperationName() string {
	return self.Identifier
}

func (self *OperationDeclarations) GetParams() ScopingInterfaces.IParameterDeclarations {
	return self.params
}

func (self *OperationDeclarations) GetExceptionList() interface{} {
	return self.exceptionList
}

func (self *OperationDeclarations) String() string {
	return fmt.Sprintf("OperationDeclarations: %v, %v, ReturnType: %v", self.TypeSpecBase.String(), self.Identifier, self.declarator.GetName())
}
