package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type SequenceType struct {
	TypeSpecBase
	declaredType ScopingInterfaces.IBaseDeclaredType
	count        int
}

func NewSequenceType(fileInformation ScopingInterfaces.IFileInformation, declaredType ScopingInterfaces.IBaseDeclaredType, count int) (*SequenceType, error) {
	identifier := fmt.Sprintf("sequence_%v", declaredType.GetName())

	return &SequenceType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.SequenceIdlType,
			false,
			false,
			false,
			false,
			false),
		declaredType: declaredType,
		count:        count,
	}, nil
}

func (self *SequenceType) TypeSpec() ScopingInterfaces.IBaseDeclaredType {
	return self.declaredType
}

func (self *SequenceType) Count() int {
	return self.count
}
