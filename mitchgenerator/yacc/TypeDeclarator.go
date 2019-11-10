package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type TypeDeclarator struct {
	DefinedTyped interfaces.IDefinedType
	Declarator   interfaces.IDeclarator
	next         interfaces.IDefinitionDeclaration
}

func (self *TypeDeclarator) GetStreamFunctionName() string {
	return "Aasdasdasas"
}

func (self *TypeDeclarator) GetPackageName() (bool, string, string) {
	return self.DefinedTyped.GetPackageName()
}

func (self *TypeDeclarator) GetSequenceCount() (bool, int) {
	return self.DefinedTyped.GetSequenceCount()
}

func (self *TypeDeclarator) Kind() interfaces.Kind {
	return self.DefinedTyped.Kind()
}

func (self *TypeDeclarator) DefaultValue() string {
	return self.DefinedTyped.DefaultValue()
}

func (self *TypeDeclarator) Predefined() bool {
	return self.DefinedTyped.Predefined()
}

func (self *TypeDeclarator) GetScopeName() string {
	return self.Declarator.Identifier()
}

func (self *TypeDeclarator) GetNext() interfaces.IDefinitionDeclaration {
	return self.next
}

func (self *TypeDeclarator) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.next = typeSpec
}

func (self *TypeDeclarator) ClearNext() {
	self.next = nil
}

func (self *TypeDeclarator) GetDefinedTyped() interfaces.IDefinedType {
	return self.DefinedTyped
}

func (self *TypeDeclarator) GetDeclarator() interfaces.IDeclarator {
	return self.Declarator
}

func (self *TypeDeclarator) GetName() string {
	return fmt.Sprintf(self.Declarator.Identifier())
}

func Newtypedef_dcl(definedTyped interfaces.IDefinedType, declarator interfaces.IDeclarator) *TypeDeclarator {
	if definedTyped == nil {
		return nil
	}
	if declarator == nil {
		return nil
	}
	return &TypeDeclarator{
		DefinedTyped: definedTyped,
		Declarator:   declarator,
	}
}
