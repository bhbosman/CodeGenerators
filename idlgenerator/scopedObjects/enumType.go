package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type EnumType struct {
	TypeSpecBase
	enumerator ScopingInterfaces.IEnumerator
}

func NewEnumType(fileInformation ScopingInterfaces.IFileInformation, id string, enumerator ScopingInterfaces.IEnumerator) (ScopingInterfaces.IEnumType, error) {
	if id == "" {
		return nil, fmt.Errorf("invalid enum enum type")
	}
	if enumerator == nil {
		return nil, fmt.Errorf("invalid enumerator")
	}

	return &EnumType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			id,
			ScopingInterfaces.RWEnumIdlType,
			false,
			false,
			false,
			false),
		enumerator: enumerator,
	}, nil
}

func (self *EnumType) String() string {
	return fmt.Sprintf("EnumType: %v, %v", self.FileInformationBase.String(), self.Identifier)
}

func (e *EnumType) Enumerator() ScopingInterfaces.IEnumerator {
	return e.enumerator
}
