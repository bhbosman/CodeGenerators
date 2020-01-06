package scopedObjects

import (
	//"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	//"go.uber.org/multierr"
)

type ValueAbsoluteDefinitionFlags uint8

const (
	VADNone                                 = 0
	VADForward ValueAbsoluteDefinitionFlags = 1 << iota
	VADAbstract
)

func NewIdlValueAbsoluteDefinition(fileInformation si.IFileInformation, identifier string, typeSpec si.ITypeSpec, flags ValueAbsoluteDefinitionFlags) (si.IInterfaceDcl, error) {

	return NewInterfaceDcl(
		fileInformation,
		si.IdlValue_Abs_DefType,
		identifier,
		flags&VADForward == VADForward,
		flags&VADAbstract == VADAbstract,
		false,
		typeSpec)
}
