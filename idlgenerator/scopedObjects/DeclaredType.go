package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type DeclaredType struct {
	TypeSpecBase
}

func NewDeclaredType(kind ScopingInterfaces.IDlSupportedTypes, fileName string, row int, col int) *DeclaredType {
	info := NewFileInformationBase01(fileName, row, col)
	return &DeclaredType{
		TypeSpecBase: NewTypeSpecBase(
			&info,
			nil,
			kind.IDLToken(),
			kind, true,
			false,
			false,
			false,
			false),
	}
}
