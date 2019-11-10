package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type Idlvalue_forward_dcl struct {
	TypeSpecBase
}

func NewIdlvalue_forward_dcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, fileName string, row int, col int) (*Idlvalue_forward_dcl, error) {
	return &Idlvalue_forward_dcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.Idlvalue_forward_dclType,
			false,
			false,
			false,
			false),
	}, nil
}
