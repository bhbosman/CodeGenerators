package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type SequenceType struct {
	TypeSpecBase
	declaredType ScopingInterfaces.IDeclaredType
	count        int
}

func NewSequenceType(fileInformation ScopingInterfaces.IFileInformation, declaredType ScopingInterfaces.IDeclaredType, count int) (*SequenceType, error) {
	identifier := fmt.Sprintf("sequence_%v_%v", declaredType.GetName(), count)
	return &SequenceType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.SequenceIdlType,
			false,
			false,
			false,
			false),
		declaredType: declaredType,
		count:        count,
	}, nil
}

func (self *SequenceType) TypeSpec() ScopingInterfaces.IDeclaredType {
	return self.declaredType
}

func (self *SequenceType) Count() int {
	return self.count
}
