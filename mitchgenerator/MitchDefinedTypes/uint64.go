package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchUInt64 struct {
}

func (self *MitchUInt64) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchUInt64) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchUInt64) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt64) DefaultValue() string {
	return "0"
}

func (self *MitchUInt64) Kind() interfaces.Kind {
	return interfaces.MitchUInt64
}

func (self *MitchUInt64) Predefined() bool {
	return true
}

func (self *MitchUInt64) GetName() string {
	return self.Kind().String()
}
