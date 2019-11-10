package interfaces

type ITypeDeclaration interface {
	IDefinitionDeclaration
	GetDefinedTyped() IDefinedType
	GetDeclarator() IDeclarator
}
