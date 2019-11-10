package scopedObjects

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type IdlValueHeader struct {
	fileInformation      ScopingInterfaces.IFileInformation
	idlValueKind         ScopingInterfaces.IIdlValueKind
	id                   ScopingInterfaces.IIdlIdentifier
	valueInheritanceSpec ScopingInterfaces.IValueInheritanceSpec
}

func NewIdlValueHeader(fileInformation ScopingInterfaces.IFileInformation, idlValueKind ScopingInterfaces.IIdlValueKind, id ScopingInterfaces.IIdlIdentifier, valueInheritanceSpec ScopingInterfaces.IValueInheritanceSpec) (ScopingInterfaces.IIdlValueHeader, error) {
	return &IdlValueHeader{
		fileInformation:      fileInformation,
		idlValueKind:         idlValueKind,
		id:                   id,
		valueInheritanceSpec: valueInheritanceSpec,
	}, nil
}

func (self IdlValueHeader) GetFileName() string {
	return self.fileInformation.GetFileName()
}

func (self IdlValueHeader) GetRow() int {
	return self.fileInformation.GetRow()
}

func (self IdlValueHeader) GetCol() int {
	return self.fileInformation.GetCol()
}
