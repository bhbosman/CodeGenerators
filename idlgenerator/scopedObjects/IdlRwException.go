package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type IdlRwException struct {
	TypeSpecBase
	members ScopingInterfaces.IStructMember
}

func (self *IdlRwException) Members() ScopingInterfaces.IStructMember {
	return self.members
}

func NewIdlRwException(fileInformation ScopingInterfaces.IFileInformation, identifier string, members ScopingInterfaces.IStructMember) (*IdlRwException, error) {
	return &IdlRwException{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.ExceptionIdlType,
			false,
			false,
			false,
			false,
			false),
		members: members,
	}, nil
}

func (self *IdlRwException) String() string {
	return fmt.Sprintf("IdlRwException: %v, %v", self.TypeSpecBase.String(), self.Identifier)
}
