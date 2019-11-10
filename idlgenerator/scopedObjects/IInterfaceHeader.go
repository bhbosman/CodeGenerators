package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type InterfaceHeader struct {
	FileInformationBase
	identifier               string
	local                    bool
	abstract                 bool
	interfaceInheritanceSpec ScopingInterfaces.IInterfaceInheritanceSpec
}

func NewInterfaceHeader(fileInformation ScopingInterfaces.IFileInformation, identifier string, local, abstract bool, interfaceInheritanceSpec ScopingInterfaces.IInterfaceInheritanceSpec) (*InterfaceHeader, error) {
	return &InterfaceHeader{
		FileInformationBase:      NewFileInformationBase02(fileInformation),
		identifier:               identifier,
		local:                    local,
		abstract:                 abstract,
		interfaceInheritanceSpec: interfaceInheritanceSpec,
	}, nil
}
func (self *InterfaceHeader) Local() bool {
	return self.local
}

func (self *InterfaceHeader) Abstract() bool {
	return self.abstract
}

func (self *InterfaceHeader) Identifier() string {
	return self.identifier
}
