package scopedObjects

import (
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type DeclaredTypePlaceHolder struct {
	FileInformationBase
	identifier string
	kind       si.IDlSupportedTypes
}

func (self *DeclaredTypePlaceHolder) AssignDeclaredTypeValues() {
	//panic("implement me")
}

func NewEmptyIdlDefinition(fileInformation si.IFileInformation, identifier string) si.IBaseDeclaredType {
	return &DeclaredTypePlaceHolder{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		identifier:          identifier,
		kind:                si.DeclareTypePlaceHolderType,
	}
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
