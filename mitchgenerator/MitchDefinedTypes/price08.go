package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchPrice08 struct {
}

func (self *MitchPrice08) GetStreamFunctionName() string {
	return "mitch_price08"
}

func (self *MitchPrice08) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchPrice08) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchPrice08) DefaultValue() string {
	return "0.0"
}

func (self *MitchPrice08) Kind() interfaces.Kind {
	return interfaces.MitchPrice08
}

func (self *MitchPrice08) Predefined() bool {
	return true
}

func (self *MitchPrice08) GetName() string {
	return self.Kind().String()
}
