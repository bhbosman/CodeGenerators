package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type StructMember struct {
	typeSpec   ScopingInterfaces.IDeclaredType
	declarator ScopingInterfaces.IDeclarator
	next       ScopingInterfaces.IStructMember
}

func NewStructMember(typeSpec ScopingInterfaces.IDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.IStructMember, error) {
	if typeSpec == nil {
		return nil, fmt.Errorf("type spec needed for new member")
	}
	if declarator == nil {
		return nil, fmt.Errorf("declarators needed for new member")
	}

	return &StructMember{
		typeSpec:   typeSpec,
		declarator: declarator,
	}, nil
}

func NewStructMember01(declaredType ScopingInterfaces.IDeclaredType, declarator ScopingInterfaces.IDeclarator, nextStructMember ScopingInterfaces.IStructMember) ScopingInterfaces.IStructMember {
	result, _ := NewStructMember(declaredType, declarator)
	result.NextStructMember(nextStructMember)
	return result
}

func (self *StructMember) GetNext() ScopingInterfaces.IStructMember {
	return self.next
}

func (self *StructMember) DeclaredType() ScopingInterfaces.IDeclaredType {
	return self.typeSpec
}

func (self *StructMember) GetDeclarator() ScopingInterfaces.IDeclarator {
	return self.declarator
}

func (self *StructMember) GetMembers() []ScopingInterfaces.IStructMemberInformation {
	buildMembers := func(member ScopingInterfaces.IStructMember) []ScopingInterfaces.IStructMemberInformation {
		result := make([]ScopingInterfaces.IStructMemberInformation, 0, 10)
		for member != nil {
			for currentDeclarator := member.GetDeclarator(); currentDeclarator != nil; currentDeclarator = currentDeclarator.GetNext() {
				fileInformation := NewFileInformationBase01(currentDeclarator.GetFileName(), currentDeclarator.GetRow(), currentDeclarator.GetCol())
				newMember := NewMemberInformation(&fileInformation, currentDeclarator.GetIdentifier(), member.DeclaredType())
				result = append(result, newMember)
			}
			member = member.GetNext()
		}
		return result
	}
	return buildMembers(self)
}

func (self *StructMember) Count() int {
	return len(self.GetMembers())
}

func (self *StructMember) NextStructMember(next ScopingInterfaces.IStructMember) ScopingInterfaces.IStructMember {
	self.next = next
	return self
}
