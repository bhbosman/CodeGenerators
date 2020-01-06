package scopedObjects

import si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type IdlValueHeader struct {
	fileInformation      si.IFileInformation
	idlValueKind         si.IIdlValueKind
	id                   si.IIdlIdentifier
	valueInheritanceSpec si.IValueInheritanceSpec
}

func NewIdlValueHeader(fileInformation si.IFileInformation, idlValueKind si.IIdlValueKind, id si.IIdlIdentifier, valueInheritanceSpec si.IValueInheritanceSpec) (si.IIdlValueHeader, error) {
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
