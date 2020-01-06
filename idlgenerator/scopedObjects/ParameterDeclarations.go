package scopedObjects

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type ParameterDeclarations struct {
	FileInformationBase
	nextParameterDeclarations si.IParameterDeclarations
	paramIn                   bool
	paramOut                  bool
	paramName                 string
	paramDeclarationType      si.IBaseDeclaredType
}

func NewParameterDeclarations(
	fileInformation si.IFileInformation,
	paramIn, paramOut bool,
	paramName string,
	paramDeclarationType si.IBaseDeclaredType) *ParameterDeclarations {

	return &ParameterDeclarations{
		FileInformationBase:       NewFileInformationBase02(fileInformation),
		nextParameterDeclarations: nil,
		paramIn:                   paramIn,
		paramOut:                  paramOut,
		paramName:                 paramName,
		paramDeclarationType:      paramDeclarationType,
	}
}

func (self *ParameterDeclarations) String() string {
	return fmt.Sprintf("ParameterDeclarations: %v, %v: %v", self.FileInformationBase.String(), self.paramName, self.paramDeclarationType.GetName())
}

func (self *ParameterDeclarations) GetNextParameterDeclarations() si.IParameterDeclarations {
	return self.nextParameterDeclarations
}

func (self *ParameterDeclarations) GetParamIn() bool {
	return self.paramIn
}

func (self *ParameterDeclarations) GetParamOut() bool {
	return self.paramOut
}

func (self *ParameterDeclarations) GetParamName() string {
	return self.paramName
}

func (self *ParameterDeclarations) GetParamDeclarationType() si.IBaseDeclaredType {
	return self.paramDeclarationType
}

func (self *ParameterDeclarations) NextParameterDeclarations(
	nextParameterDeclarations si.IParameterDeclarations) si.IParameterDeclarations {
	self.nextParameterDeclarations = nextParameterDeclarations
	return self
}
