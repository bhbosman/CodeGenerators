package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchUInt16 struct {
}

func (self *MitchUInt16) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchUInt16) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchUInt16) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt16) DefaultValue() string {
	return "0"
}

func (self *MitchUInt16) Kind() interfaces.Kind {
	return interfaces.MitchUInt16
}

func (self *MitchUInt16) Predefined() bool {
	return true
}

func (self *MitchUInt16) GetName() string {
	return self.Kind().String()
}
