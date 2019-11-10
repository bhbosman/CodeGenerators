package MitchDefinedTypes

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchAlpha struct {
	Length int64
}

func (self *MitchAlpha) GetStreamFunctionName() string {
	return self.Kind().String()
}

func (self *MitchAlpha) GetPackageName() (bool, string, string) {
	return true, "", self.Kind().String()
}

func (self *MitchAlpha) GetSequenceCount() (bool, int) {
	if self.Length == 0 {
		return false, 0
	}
	return true, int(self.Length)
}

func (self *MitchAlpha) DefaultValue() string {
	return "\"\""
}

func (self *MitchAlpha) Kind() interfaces.Kind {
	return interfaces.MitchAlpha
}

func (self *MitchAlpha) Predefined() bool {
	return true
}

func (self *MitchAlpha) GetName() string {
	return self.Kind().String()
}
