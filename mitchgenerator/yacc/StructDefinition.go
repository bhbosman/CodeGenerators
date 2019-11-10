package yacc

import (
	"encoding/json"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type MitchMessageDefinition struct {
	Type       string                            `json:"Type"`
	Identifier string                            `json:"Identifier"`
	Members    []*Member                         `json:"Members"`
	Next       interfaces.IDefinitionDeclaration `json:"-"`

	MessageLength uint16
	MessageType   byte
}

func (self *MitchMessageDefinition) HasMessageInformation() bool {
	return self.MessageType != 0 && self.MessageLength != 0
}

func (self *MitchMessageDefinition) GetStreamFunctionName() string {
	return self.Identifier
}

func (self *MitchMessageDefinition) GetPackageName() (bool, string, string) {
	return true, "", self.Identifier
}

func (self *MitchMessageDefinition) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchMessageDefinition) Kind() interfaces.Kind {
	return interfaces.Struct
}

func (self *MitchMessageDefinition) DefaultValue() string {
	return "nil"
}

func (self *MitchMessageDefinition) Predefined() bool {
	return false
}

func (self *MitchMessageDefinition) GetScopeName() string {
	return self.Identifier
}

func (self *MitchMessageDefinition) ClearNext() {
	self.Next = nil
}

func (self *MitchMessageDefinition) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *MitchMessageDefinition) GetName() string {
	return self.Identifier

}

func (self *MitchMessageDefinition) AddMember(typeSpec interfaces.IDefinedType, declarator interfaces.IDeclarator) {
	self.Members = append(self.Members, NewMember(typeSpec, declarator, nil))
}

func (self *MitchMessageDefinition) String() string {
	bytes, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func (self *MitchMessageDefinition) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func NewMitchMessageDefinition(identifier string, MessageLength int64, MessageType int64) *MitchMessageDefinition {
	return &MitchMessageDefinition{
		Type:          "struct",
		Identifier:    identifier,
		Members:       make([]*Member, 0),
		Next:          nil,
		MessageLength: uint16(MessageLength),
		MessageType:   byte(MessageType),
	}
}
