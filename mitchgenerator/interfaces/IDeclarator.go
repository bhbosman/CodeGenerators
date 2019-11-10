package interfaces

type IDeclarator interface {
	Identifier() string
	Next() IDeclarator
	SetNext(next IDeclarator)
	GetNext() IDeclarator
	ClearNext()
	DefaultValue() IConstantValue
}
