package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type FixedPointType struct {
	TypeSpecBase
	n1, n2 int
}

func NewFixedPointType(fileInformation ScopingInterfaces.IFileInformation, n1, n2 int) (ScopingInterfaces.ITypeSpec, error) {
	identifier := fmt.Sprintf("fixed_%v_%v", n1, n2)
	return &FixedPointType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.FixedPointTypeIdlType,
			false,
			false,
			false,
			false),
		n1: n1,
		n2: n2,
	}, nil
}
