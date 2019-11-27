package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type IdlConstDcl struct {
	TypeSpecBase
	value int
}

func NewIdlConstDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, value int) *IdlConstDcl {
	return &IdlConstDcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.ConstDclType,
			false,
			false,
			false,
			false,
			false),
		value: value,
	}
}

func (self *IdlConstDcl) String() string {
	return fmt.Sprintf("IdlConstDcl: %v, %v = %v", self.TypeSpecBase.String(), self.Identifier, self.value)
}
