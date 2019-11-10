package MitchDefinedTypes

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type MitchTime struct {
}

func (self *MitchTime) GetStreamFunctionName() string {
	return "mitch_time"
}

func (self *MitchTime) GetPackageName() (bool, string, string) {
	return true, "time", self.Kind().String()
}

func (self *MitchTime) GetSequenceCount() (bool, int) {
	return false, 0
}

func (self *MitchTime) DefaultValue() string {
	_, s, _ := self.GetPackageName()
	name := self.GetName()
	return fmt.Sprintf("%v.%v{}", s, name)
}

func (self *MitchTime) Kind() interfaces.Kind {
	return interfaces.MitchTime
}

func (self *MitchTime) Predefined() bool {
	return true
}

func (self *MitchTime) GetName() string {
	return "Time"
}
