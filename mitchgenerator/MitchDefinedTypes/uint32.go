package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchUInt32 struct {
}

func (self *MitchUInt32) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchUInt32) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchUInt32) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt32) DefaultValue() string {
	return "0"
}

func (self *MitchUInt32) Kind() interfaces.Kind {
	return interfaces.MitchUInt32
}

func (self *MitchUInt32) Predefined() bool {
	return true
}

func (self *MitchUInt32) GetName() string {
	return self.Kind().String()
}
