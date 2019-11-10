package yacc

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type Declarator struct {
	identifier   string
	defaultValue interfaces.IConstantValue
	next         interfaces.IDeclarator `json:"-"`
}

func (self *Declarator) DefaultValue() interfaces.IConstantValue {
	return self.defaultValue
}

func (self *Declarator) ClearNext() {
	self.next = nil
}

func (self *Declarator) GetNext() interfaces.IDeclarator {
	return self.next
}

func (self *Declarator) Identifier() string {
	return self.identifier
}

func (self *Declarator) Next() interfaces.IDeclarator {
	return self.next
}

func (self *Declarator) SetNext(next interfaces.IDeclarator) {
	self.next = next
}

func (receiver *Declarator) String() string {
	return fmt.Sprintf("Declarator => GetName: %v\n", receiver.Identifier())
}

func (self *Declarator) MarshalJSON() ([]byte, error) {

	if self.defaultValue == nil {
		return json.Marshal(&struct {
			Type       string `json:"Type"`
			Identifier string `json:"Identifier"`
		}{
			Type:       "Declarator",
			Identifier: self.identifier,
		})
	}

	return json.Marshal(&struct {
		Type                  string      `json:"Type"`
		Identifier            string      `json:"Identifier"`
		DefaultValue          interface{} `json:"DefaultValue"`
		DefaultValueType      string      `json:"DefaultValueType"`
		DefaultValueMaxLength int         `json:"DefaultValueMaxLength"`
	}{
		Type:                  "Declarator",
		Identifier:            self.identifier,
		DefaultValue:          self.defaultValue.Value(),
		DefaultValueType:      self.defaultValue.ValueKind().String(),
		DefaultValueMaxLength: self.defaultValue.MaxLength(),
	})
}

func NewDeclarator(identifier string, defaultValue interfaces.IConstantValue) *Declarator {
	return &Declarator{
		identifier:   identifier,
		defaultValue: defaultValue,
		next:         nil,
	}
}
