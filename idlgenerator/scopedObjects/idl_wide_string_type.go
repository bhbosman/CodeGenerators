package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type WideStringType struct {
	TypeSpecBase
}

func NewWideStringType(fileInformation ScopingInterfaces.IFileInformation, n int) (ScopingInterfaces.ITypeSpec, error) {
	id := fmt.Sprintf("string_")
	return &WideStringType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			id,
			ScopingInterfaces.StringIdlType,
			false,
			false,
			false,
			false),
	}, nil

}
