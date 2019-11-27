package scopedObjects

import (
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type TypeSpecBase struct {
	FileInformationBase
	NextTypeSpec     si.ITypeSpec
	Identifier       string
	Kind             si.IDlSupportedTypes
	isPrimitive      bool
	forward          bool
	abstract         bool
	local            bool
	linkedItems      []si.IBaseDeclaredType
	sequenceRequired bool
}

func NewTypeSpecBase(
	fileInformation si.IFileInformation,
	nextTypeSpec si.ITypeSpec,
	identifier string,
	kind si.IDlSupportedTypes,
	isPrimitive bool,
	forward bool,
	abstract bool,
	local bool,
	sequenceRequired bool) TypeSpecBase {
	return TypeSpecBase{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		NextTypeSpec:        nextTypeSpec,
		Identifier:          identifier,
		Kind:                kind,
		isPrimitive:         isPrimitive,
		forward:             forward,
		abstract:            abstract,
		local:               local,
		linkedItems:         make([]si.IBaseDeclaredType, 0, 16),
		sequenceRequired:    sequenceRequired,
	}
}

func (self *TypeSpecBase) Local() bool {
	return self.local
}

func (self *TypeSpecBase) GetLinkedItems() []si.IBaseDeclaredType {
	return self.linkedItems
}

func (self *TypeSpecBase) Link(placeHolder si.IDeclaredTypePlaceHolder) error {
	if placeHolder != nil {
		self.linkedItems = append(self.linkedItems, placeHolder)
		placeHolder.SetName(self.Identifier)
		placeHolder.AssignDeclaredTypeValues()
	}
	return nil
}

func (self *TypeSpecBase) UsageCount() int {
	return len(self.linkedItems)
}

func (self *TypeSpecBase) SetNextTypeSpec(next si.ITypeSpec) error {
	return self.AssignNextTypeSpec(next)
}

func (self *TypeSpecBase) AssignNextTypeSpec(next si.ITypeSpec) error {
	if self.NextTypeSpec == nil {
		self.NextTypeSpec = next
	} else {
		return self.NextTypeSpec.SetNextTypeSpec(next)
	}
	return nil
}

func (self *TypeSpecBase) GetKind() si.IDlSupportedTypes {
	return self.Kind
}

func (self *TypeSpecBase) GetNextTypeSpec() (si.ITypeSpec, error) {
	return self.NextTypeSpec, nil
}

func (self *TypeSpecBase) GetName() string {
	return self.Identifier
}

func (self *TypeSpecBase) SetName(name string) {
	self.Identifier = name
	for _, declType := range self.linkedItems {
		declType.SetName(self.Identifier)
	}
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
