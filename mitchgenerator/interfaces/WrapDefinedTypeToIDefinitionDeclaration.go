package interfaces

type WrapDefinedTypeToIDefinitionDeclaration struct {
	inner IDefinedType
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetStreamFunctionName() string {
	return self.inner.GetStreamFunctionName()
}

func NewWrapDefinedTypeToIDefinitionDeclaration(inner IDefinedType) *WrapDefinedTypeToIDefinitionDeclaration {
	return &WrapDefinedTypeToIDefinitionDeclaration{inner: inner}
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetPackageName() (bool, string, string) {
	return self.inner.GetPackageName()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetSequenceCount() (bool, int) {
	return self.inner.GetSequenceCount()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetName() string {
	return self.inner.GetName()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) Kind() Kind {
	return self.inner.Kind()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) DefaultValue() string {
	return self.inner.DefaultValue()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) Predefined() bool {
	return self.inner.Predefined()
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetNext() IDefinitionDeclaration {
	return nil
}

func (self *WrapDefinedTypeToIDefinitionDeclaration) SetNext(typeSpec IDefinitionDeclaration) {

}

func (self *WrapDefinedTypeToIDefinitionDeclaration) ClearNext() {

}

func (self *WrapDefinedTypeToIDefinitionDeclaration) GetScopeName() string {
	return self.inner.GetName()
}
