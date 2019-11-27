package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type InterfaceDcl struct {
	TypeSpecBase
	body ScopingInterfaces.ITypeSpec
}

func NewInterfaceDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, forward, abstract, local bool, body ScopingInterfaces.ITypeSpec) (*InterfaceDcl, error) {
	return &InterfaceDcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			identifier,
			ScopingInterfaces.InterfaceIdlType,
			false,
			forward,
			abstract,
			local,
			false),
		body: func() ScopingInterfaces.ITypeSpec {
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

func (self *InterfaceDcl) Iterate(cb func(typeSpec ScopingInterfaces.ITypeSpec) error) error {
	var err error
	if cb != nil && self.body != nil {
		for d := self.body; d != nil; d, _ = d.GetNextTypeSpec() {
			typeSpec, ok := d.(ScopingInterfaces.ITypeSpec)
			if ok {
				err = multierr.Append(err, cb(typeSpec))
			}
		}
	}
	return err
}

func (self *InterfaceDcl) GetBody() ScopingInterfaces.ITypeSpec {
	return self.body
}

func (self *InterfaceDcl) BodyArray() []ScopingInterfaces.IIdlDefinition {
	result := make([]ScopingInterfaces.IIdlDefinition, 0, 16)
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
