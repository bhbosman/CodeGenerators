package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type IdlModuleDcl struct {
	TypeSpecBase
	moduleExports ScopingInterfaces.ITypeSpec
}

func (self *IdlModuleDcl) IsDefined() bool {
	panic("implement me")
}

func (self *IdlModuleDcl) SetModuleExports(moduleExports ScopingInterfaces.ITypeSpec) {
	self.moduleExports = moduleExports
}

func (self *IdlModuleDcl) Iterate(cb func(TypeSpec ScopingInterfaces.ITypeSpec) error) error {
	var err error
	if cb != nil && self.moduleExports != nil {
		for item := self.moduleExports; item != nil; item, _ = item.GetNextTypeSpec() {
			err = multierr.Append(err, cb(item))
		}
	}
	return err
}

func (self *IdlModuleDcl) ModuleDcl() ScopingInterfaces.IIdlModuleDcl {
	return self
}

func NewModuleDcl(fileInformation ScopingInterfaces.IFileInformation, moduleIdentifier string, moduleExports ScopingInterfaces.ITypeSpec) *IdlModuleDcl {
	return &IdlModuleDcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			moduleIdentifier,
			ScopingInterfaces.ModuleIdlType,
			false,
			false,
			false,
			false),
		moduleExports: moduleExports,
	}
}

func (self *IdlModuleDcl) GetModuleExports() ScopingInterfaces.ITypeSpec {
	return self.moduleExports
}

func (self *IdlModuleDcl) String() string {
	return fmt.Sprintf("Module: GetName %v", self.Identifier)
}
