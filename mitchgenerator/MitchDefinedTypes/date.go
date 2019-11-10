package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchDate struct {
}

func (self *MitchDate) GetStreamFunctionName() string {
	return "mitch_date"
}

func (self *MitchDate) GetPackageName() (bool, string, string) {
	return true, "time", self.Kind().String()
}

func (self *MitchDate) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchDate) DefaultValue() string {
	_, s, _ := self.GetPackageName()
	name := self.GetName()
	return fmt.Sprintf("%v.%v{}", s, name)
}

func (self *MitchDate) Kind() interfaces.Kind {
	return interfaces.MitchDate
}

func (self *MitchDate) Predefined() bool {
	return true
}

func (self *MitchDate) GetName() string {
	return "Time"
}
