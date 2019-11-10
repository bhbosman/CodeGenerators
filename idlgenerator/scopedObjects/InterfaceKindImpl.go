package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type InterfaceKindImpl struct {
	FileInformationBase
	local    bool
	abstract bool
}

func NewInterfaceKindImpl(fileInformation ScopingInterfaces.IFileInformation, local bool, abstract bool) *InterfaceKindImpl {
	return &InterfaceKindImpl{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		local:               local,
		abstract:            abstract,
	}
}

func (self *InterfaceKindImpl) String() string {
	return fmt.Sprintf("InterfaceKindImpl: %v, local: %v, abstract: %v", self.FileInformationBase.String(), self.local, self.abstract)
}

func (i InterfaceKindImpl) Local() bool {
	return i.local
}

func (i InterfaceKindImpl) Abstract() bool {
	return i.abstract
}
