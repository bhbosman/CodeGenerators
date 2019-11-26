package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

//go:generate goyacc -o completeIdl.go  -p "CompleteIdl"  completeIdl.y

type StructType struct {
	TypeSpecBase
	members ScopingInterfaces.IStructMember
}

func NewStructType(fileInformation ScopingInterfaces.IFileInformation, identifier string, members ScopingInterfaces.IStructMember, forward bool) *StructType {
	return &StructType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.StructIdlType,
			false,
			forward,
			false,
			false),
		members: members,
	}
}

func (self *StructType) FindMemberType(memberIdentifier string) ScopingInterfaces.IBaseDeclaredType {
	if self.members == nil {
		return nil
	}
	for member := self.members; member != nil; member = member.GetNext() {
		for declarator := member.GetDeclarator(); declarator != nil; declarator = declarator.GetNext() {
			if memberIdentifier == declarator.GetIdentifier() {
				return member.DeclaredType()
			}
		}
	}
	return nil
}

func (self *StructType) Members() ScopingInterfaces.IStructMember {
	return self.members
}

func (self *StructType) Create() ScopingInterfaces.IIdlComparer {
	return &StructTypeComparer{}
}

func (self *StructType) String() string {
	return fmt.Sprintf("StructType: %v, %v, forward: %v", self.TypeSpecBase.String(), self.Identifier, self.forward)
}
