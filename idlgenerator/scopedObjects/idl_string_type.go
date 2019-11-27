package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type StringType struct {
	TypeSpecBase
}

func NewStringType(fileInformation ScopingInterfaces.IFileInformation, n int) (ScopingInterfaces.ITypeSpec, error) {
	identifier := func(n int) string {
		if n == 0 {
			return fmt.Sprintf("string")
		}
		return fmt.Sprintf("string_%v", n)
	}(n)
	return &StringType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.AnsiStringIdlType,
			true,
			false,
			false,
			false,
			false),
	}, nil
}
