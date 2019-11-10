package MitchDefinedTypes

import (
	"encoding/json"
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchBitField struct {
	Identifier string
	BitField00 string
	BitField01 string
	BitField02 string
	BitField03 string
	BitField04 string
	BitField05 string
	BitField06 string
	BitField07 string
	BitsUsed   byte
	Next       interfaces.IDefinitionDeclaration
}

func (self *MitchBitField) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchBitField) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchBitField) GetPackageName() (bool, string, string) {
	return false, "", self.GetName()
}

func (self *MitchBitField) Kind() interfaces.Kind {
	return interfaces.MitchBitField
}

func (self *MitchBitField) DefaultValue() string {
	return fmt.Sprintf("%v{}", self.Identifier)
}

func (self *MitchBitField) Predefined() bool {
	return false
}

var MitchBitFieldCounter = 0

func getNumber() int {
	MitchBitFieldCounter += 1
	return MitchBitFieldCounter
}

func (self *MitchBitField) GetName() string {
	return self.Identifier
}

func (self *MitchBitField) GetNext() interfaces.IDefinitionDeclaration {
	return self.Next
}

func (self *MitchBitField) SetNext(typeSpec interfaces.IDefinitionDeclaration) {
	self.Next = typeSpec
}

func (self *MitchBitField) ClearNext() {
	self.Next = nil
}

func (self *MitchBitField) GetScopeName() string {
	return self.GetName()
}

func (self *MitchBitField) MarshalJSON() ([]byte, error) {

	return json.Marshal(&struct {
		Type            string `json:"Type"`
		Identifier      string `json:"Identifier"`
		MitchBitField00 string `json:"MitchBitField00"`
		MitchBitField01 string `json:"MitchBitField01"`
		MitchBitField02 string `json:"MitchBitField02"`
		MitchBitField03 string `json:"MitchBitField03"`
		MitchBitField04 string `json:"MitchBitField04"`
		MitchBitField05 string `json:"MitchBitField05"`
		MitchBitField06 string `json:"MitchBitField06"`
		MitchBitField07 string `json:"MitchBitField07"`
		BitsUsed        byte   `json:"BitsUsed"`
	}{
		Type:            "MitchBitField",
		Identifier:      self.Identifier,
		MitchBitField00: self.BitField00,
		MitchBitField01: self.BitField01,
		MitchBitField02: self.BitField02,
		MitchBitField03: self.BitField03,
		MitchBitField04: self.BitField04,
		MitchBitField05: self.BitField05,
		MitchBitField06: self.BitField06,
		MitchBitField07: self.BitField07,
		BitsUsed:        self.BitsUsed,
	})
}

func NewMitchBitField(
	MitchBitField00 string,
	MitchBitField01 string,
	MitchBitField02 string,
	MitchBitField03 string,
	MitchBitField04 string,
	MitchBitField05 string,
	MitchBitField06 string,
	MitchBitField07 string) *MitchBitField {

	used := byte(0)
	if MitchBitField00 != "b0" {
		used = used | 0x01
	}
	if MitchBitField01 != "b1" {
		used = used | 0x02
	}
	if MitchBitField02 != "b2" {
		used = used | 0x04
	}
	if MitchBitField03 != "b3" {
		used = used | 0x08
	}
	if MitchBitField04 != "b4" {
		used = used | 0x10
	}
	if MitchBitField05 != "b5" {
		used = used | 0x20
	}
	if MitchBitField06 != "b6" {
		used = used | 0x40
	}
	if MitchBitField07 != "b7" {
		used = used | 0x80
	}

	number := getNumber()

	return &MitchBitField{
		Identifier: fmt.Sprintf("MitchBitField%04d", number),
		BitField00: MitchBitField00,
		BitField01: MitchBitField01,
		BitField02: MitchBitField02,
		BitField03: MitchBitField03,
		BitField04: MitchBitField04,
		BitField05: MitchBitField05,
		BitField06: MitchBitField06,
		BitField07: MitchBitField07,
		BitsUsed:   used,
		Next:       nil,
	}
}
