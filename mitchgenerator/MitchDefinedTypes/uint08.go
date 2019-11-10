package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchUInt08 struct {
}

func (self *MitchUInt08) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchUInt08) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchUInt08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchUInt08) DefaultValue() string {
	return "0"
}

func (self *MitchUInt08) Kind() interfaces.Kind {
	return interfaces.MitchUInt08
}

func (self *MitchUInt08) Predefined() bool {
	return true
}

func (self *MitchUInt08) GetName() string {
	return self.Kind().String()
}
