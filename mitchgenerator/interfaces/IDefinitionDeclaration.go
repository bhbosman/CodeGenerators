package interfaces

type IDefinitionDeclaration interface {
	IDefinedType
	GetNext() IDefinitionDeclaration
	SetNext(typeSpec IDefinitionDeclaration)
	ClearNext()
	GetScopeName() string
}
