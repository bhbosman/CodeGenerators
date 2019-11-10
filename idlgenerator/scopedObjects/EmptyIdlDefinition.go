package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type EmptyIdlDefinition struct {
	FileInformationBase
	identifier string
	kind       ScopingInterfaces.IDlSupportedTypes
}

func (self *EmptyIdlDefinition) IsPrimitive() bool {
	return false
}

func NewEmptyIdlDefinition(fileInformation ScopingInterfaces.IFileInformation, identifier string) ScopingInterfaces.IDeclaredType {
	return &EmptyIdlDefinition{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		identifier:          identifier,
		kind:                ScopingInterfaces.Unknown,
	}
}

func (self *EmptyIdlDefinition) IsDefined() bool {
	return false
}

func (self *EmptyIdlDefinition) Create() ScopingInterfaces.IIdlComparer {
	return &EmptyIdlDefinitionComparer{}
}

func (self *EmptyIdlDefinition) GetName() string {
	return self.identifier
}

func (self *EmptyIdlDefinition) SetName(name string)  {
	self.identifier = name
}


func (self *EmptyIdlDefinition) GetKind() ScopingInterfaces.IDlSupportedTypes {
	return self.kind
}
