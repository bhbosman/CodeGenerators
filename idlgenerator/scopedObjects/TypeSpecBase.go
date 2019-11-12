package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type TypeSpecBase struct {
	FileInformationBase
	NextTypeSpec ScopingInterfaces.ITypeSpec
	Identifier   string
	Kind         ScopingInterfaces.IDlSupportedTypes
	isPrimitive  bool
	forward      bool
	abstract     bool
	local        bool
}

func NewTypeSpecBase(
	fileInformation ScopingInterfaces.IFileInformation,
	nextTypeSpec ScopingInterfaces.ITypeSpec,
	identifier string,
	kind ScopingInterfaces.IDlSupportedTypes,
	isPrimitive bool,
	forward bool,
	abstract bool,
	local bool) TypeSpecBase {
	return TypeSpecBase{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		NextTypeSpec:        nextTypeSpec,
		Identifier:          identifier,
		Kind:                kind,
		isPrimitive:         isPrimitive,
		forward:             forward,
		abstract:            abstract,
		local:               local,
	}
}

func (self *TypeSpecBase) SetNextTypeSpec(next ScopingInterfaces.ITypeSpec) error {
	return self.AssignNextTypeSpec(next)
}

func (self *TypeSpecBase) AssignNextTypeSpec(next ScopingInterfaces.ITypeSpec) error {
	if self.NextTypeSpec == nil {
		self.NextTypeSpec = next
	} else {
		return self.NextTypeSpec.SetNextTypeSpec(next)
	}
	return nil
}

func (self *TypeSpecBase) GetKind() ScopingInterfaces.IDlSupportedTypes {
	return self.Kind
}

func (self *TypeSpecBase) GetNextTypeSpec() (ScopingInterfaces.ITypeSpec, error) {
	return self.NextTypeSpec, nil
}

func (self *TypeSpecBase) GetName() string {
	return self.Identifier
}

func (self *TypeSpecBase) SetName(name string) {
	self.Identifier = name
}

func (self *TypeSpecBase) Forward() bool {
	return self.forward
}

func (self *TypeSpecBase) Abstract() bool {
	return self.abstract
}

func (self *TypeSpecBase) IsPrimitive() bool {
	return self.isPrimitive
}

func (self *TypeSpecBase) IsDefined() bool {
	return true
}
