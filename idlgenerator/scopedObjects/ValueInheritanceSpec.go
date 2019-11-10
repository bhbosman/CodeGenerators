package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type ValueInheritanceSpec struct {
	FileInformationBase
}

func NewValueInheritanceSpec(fileInformation ScopingInterfaces.IFileInformation) (ScopingInterfaces.IValueInheritanceSpec, error) {
	return &ValueInheritanceSpec{
		FileInformationBase: NewFileInformationBase02(fileInformation),
	}, nil
}
