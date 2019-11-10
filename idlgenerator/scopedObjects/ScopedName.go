package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type ScopedName struct {
	IdlIdentifier
	nextScopedName ScopingInterfaces.IScopedName
}

func (self *ScopedName) SetName(string) {
	panic("implement me")
}

func NewScopedName(name string, fileName string, row int, col int) *ScopedName {
	return &ScopedName{
		IdlIdentifier:  *NewIdlIdentifier(name, fileName, row, col),
		nextScopedName: nil,
	}
}

func NewScopedName02(fileInformation ScopingInterfaces.IFileInformation, name string) *ScopedName {
	return &ScopedName{
		IdlIdentifier:  *NewIdlIdentifier02(fileInformation, name),
		nextScopedName: nil,
	}
}

func (self *ScopedName) NextScopedName(next ScopingInterfaces.IScopedName) error {
	if self.nextScopedName == nil {
		self.nextScopedName = next
	} else {
		return self.nextScopedName.NextScopedName(next)
	}
	return nil
}

func (self *ScopedName) GetNextScopedName() (ScopingInterfaces.IScopedName, error) {
	return self.nextScopedName, nil
}
