package scopedObjects

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type ValueAbsoluteDefinitionFlags uint8

const (
	VADNone                                 = 0
	VADForward ValueAbsoluteDefinitionFlags = 1 << iota
	VADAbstract
)

type ValueAbsoluteDefinition struct {
	TypeSpecBase
	typeSpec si.ITypeSpec
}

func NewIdlValueAbsoluteDefinition(fileInformation si.IFileInformation, identifier string, typeSpec si.ITypeSpec, flags ValueAbsoluteDefinitionFlags) (*ValueAbsoluteDefinition, error) {
	return &ValueAbsoluteDefinition{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			si.IdlValue_Abs_DefType,
			false,
			flags&VADForward == VADForward,
			flags&VADAbstract == VADAbstract,
			false,
			false),
		typeSpec: typeSpec,
	}, nil
}

func (self *ValueAbsoluteDefinition) Iterate(cb func(typeSpec si.ITypeSpec) error) error {
	var err error
	if cb != nil && self.typeSpec != nil {
		for b := self.typeSpec; b != nil; b, _ = b.GetNextTypeSpec() {
			err = multierr.Append(err, cb(b))
		}
	}
	return err
}

func (self *ValueAbsoluteDefinition) String() string {
	return fmt.Sprintf("ValueAbsoluteDefinition: %v, %v", self.TypeSpecBase.String(), self.Identifier)
}
