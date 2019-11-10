package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type ParamAttribute struct {
	FileInformationBase
	in  bool
	out bool
}

func (self *ParamAttribute) In() bool {
	return self.in
}

func (self *ParamAttribute) Out() bool {
	return self.out
}

func NewParamAttribute(fileInformation ScopingInterfaces.IFileInformation, in bool, out bool) *ParamAttribute {
	return &ParamAttribute{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		in:                  in,
		out:                 out,
	}
}
