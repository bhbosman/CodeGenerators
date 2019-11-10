package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchPrice04 struct {
}

func (self *MitchPrice04) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchPrice04) GetStreamFunctionName() string {
	return "mitch_price04"
}

func (self *MitchPrice04) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchPrice04) DefaultValue() string {
	return "0.0"
}

func (self *MitchPrice04) Kind() interfaces.Kind {
	return interfaces.MitchPrice04
}

func (self *MitchPrice04) Predefined() bool {
	return true
}

func (self *MitchPrice04) GetName() string {
	return self.Kind().String()
}
