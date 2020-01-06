package scopedObjects

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type InterfaceDcl struct {
	TypeSpecBase
	body si.ITypeSpec
}

func NewInterfaceDcl(
	fileInformation si.IFileInformation,
	kind si.IDlSupportedTypes,
	identifier string,
	forward bool,
	abstract bool,
	local bool,
	body si.ITypeSpec) (si.IInterfaceDcl, error) {

	return &InterfaceDcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			kind,
			false,
			forward,
			abstract,
			local,
			false),
		body: func() si.ITypeSpec {
			if body == nil {
				return nil
			}
			return body
		}(),
	}, nil
}

func (self *InterfaceDcl) Local() bool {
	return self.local
}

func (self *InterfaceDcl) Iterate(cb func(typeSpec si.ITypeSpec) error) error {
	var err error
	if cb != nil && self.body != nil {
		for d := self.body; d != nil; d, _ = d.GetNextTypeSpec() {
			typeSpec, ok := d.(si.ITypeSpec)
			if ok {
				err = multierr.Append(err, cb(typeSpec))
			}
		}
	}
	return err
}

func (self *InterfaceDcl) GetBody() si.ITypeSpec {
	return self.body
}

func (self *InterfaceDcl) BodyArray() []si.IIdlDefinition {
	result := make([]si.IIdlDefinition, 0, 16)
	for item := self.GetBody(); item != nil; item, _ = item.GetNextTypeSpec() {
		result = append(result, item)
	}
	return result
}

func (self *InterfaceDcl) BodyCount() int {
	n := 0
	for item := self.GetBody(); item != nil; item, _ = item.GetNextTypeSpec() {
		n++
	}
	return n
}

func (self *InterfaceDcl) String() string {
	return fmt.Sprintf("InterfaceDcl: %v, %v, forward: %v, abstract: %v, local: %v", self.TypeSpecBase.String(), self.Identifier, self.forward, self.abstract, self.local)
}
