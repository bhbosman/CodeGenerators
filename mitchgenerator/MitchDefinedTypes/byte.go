package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchByte struct {
}

func (self *MitchByte) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchByte) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchByte) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchByte) DefaultValue() string {
	return "0"
}

func (self *MitchByte) Kind() interfaces.Kind {
	return interfaces.MitchByte
}

func (self *MitchByte) Predefined() bool {
	return true
}

func (self *MitchByte) GetName() string {
	return self.Kind().String()
}
