package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type IdlValueKind struct {
	FileInformationBase
	custom bool
}

func NewIdlValueKind(fileInformation ScopingInterfaces.IFileInformation, custom bool) (*IdlValueKind, error) {
	return &IdlValueKind{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		custom:              custom,
	}, nil
}
