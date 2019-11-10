package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type IdlIdentifier struct {
	FileInformationBase
	name string
}

func NewIdlIdentifier(name string, fileName string, row int, col int) *IdlIdentifier {
	return &IdlIdentifier{
		FileInformationBase: NewFileInformationBase01(fileName, row, col),
		name:                name,
	}
}

func NewIdlIdentifier02(fileInformation ScopingInterfaces.IFileInformation, name string) *IdlIdentifier {
	return NewIdlIdentifier(name, fileInformation.GetFileName(), fileInformation.GetRow(), fileInformation.GetCol())
}

func (self *IdlIdentifier) GetName() string {
	return self.name
}

func (self *IdlIdentifier) Identifier() string {
	return self.name
}

func (self *IdlIdentifier) SetName(string) {
	panic("implement me")
}
