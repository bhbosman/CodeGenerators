package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type TypePrefixDefinition struct {
	TypeSpecBase
	stringLiteral string
}

func (self *TypePrefixDefinition) Create() ScopingInterfaces.IIdlComparer {
	return &TypePrefixDefinitionComparer{}
}

func NewCreateTypePrefixDcl(fileInformation ScopingInterfaces.IFileInformation, scopedName string, stringLiteral string) (*TypePrefixDefinition, error) {
	return &TypePrefixDefinition{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			scopedName,
			ScopingInterfaces.TypePrefixDefinitionIdlType,
			false,
			false,
			false,
			false),
		stringLiteral: stringLiteral,
	}, nil
}

func (self *TypePrefixDefinition) StringLiteral() string {
	return self.stringLiteral
}

func (self *TypePrefixDefinition) String() string {
	return fmt.Sprintf("Type: TypePrefixDefinition, (%v: %v)", self.Identifier, self.stringLiteral)
}
