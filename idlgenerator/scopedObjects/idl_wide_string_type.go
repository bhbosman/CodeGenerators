package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type WideStringType struct {
	TypeSpecBase
}

func NewWideStringType(fileInformation ScopingInterfaces.IFileInformation, n int) (ScopingInterfaces.ITypeSpec, error) {
	id := func() string {
		if n == 0 {
			return fmt.Sprintf("string")
		}
		return fmt.Sprintf("string_%v", n)
	}()
	return &WideStringType{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			id,
			ScopingInterfaces.WideStringIdlType,
			true,
			false,
			false,
			false,
			false),
	}, nil

}
