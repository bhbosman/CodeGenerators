package scopedObjects

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

//type IDeclareTypeInformation interface {
//
//
//}
//
//type DeclareTypeInformation struct {
//
//}

type DeclaredTypePlaceHolder struct {
	FileInformationBase
	identifier string
	kind       si.IDlSupportedTypes
}

func NewEmptyIdlDefinition(fileInformation si.IFileInformation, identifier string) si.IBaseDeclaredType {
	return &DeclaredTypePlaceHolder{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		identifier:          identifier,
		kind:                si.DeclareTypePlaceHolderType,
	}
}

func (self *DeclaredTypePlaceHolder) Link(declaredType si.IBaseDeclaredType) error {
	return fmt.Errorf("error as this is a place holder for %v", self.identifier)
}

func (self *DeclaredTypePlaceHolder) UsageCount() int {
	return -1
}

func (self *DeclaredTypePlaceHolder) IsPrimitive() bool {
	return false
}

func (self *DeclaredTypePlaceHolder) IsDefined() bool {
	return false
}

func (self *DeclaredTypePlaceHolder) Create() si.IIdlComparer {
	return &EmptyIdlDefinitionComparer{}
}

func (self *DeclaredTypePlaceHolder) GetName() string {
	return self.identifier
}

func (self *DeclaredTypePlaceHolder) SetName(name string) {
	self.identifier = name
}

func (self *DeclaredTypePlaceHolder) GetKind() si.IDlSupportedTypes {
	return self.kind
}
