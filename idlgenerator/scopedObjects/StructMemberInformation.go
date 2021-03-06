package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type StructMemberInformation struct {
	FileInformationBase
	id       string
	typeSpec ScopingInterfaces.IBaseDeclaredType
}

func NewMemberInformation(fileInformation ScopingInterfaces.IFileInformation, id string, typeSpec ScopingInterfaces.IBaseDeclaredType) *StructMemberInformation {
	return &StructMemberInformation{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		id:                  id,
		typeSpec:            typeSpec,
	}
}

func (self *StructMemberInformation) GetId() string {
	return self.id
}

func (self *StructMemberInformation) GetTypeSpec() ScopingInterfaces.IBaseDeclaredType {
	return self.typeSpec
}
func (self *StructMemberInformation) String() string {
	return fmt.Sprintf("StructMemberInformation: %v, %v: %v", self.FileInformationBase.String(), self.id, self.typeSpec.GetName())
}
