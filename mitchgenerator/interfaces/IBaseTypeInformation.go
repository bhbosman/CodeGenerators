package interfaces

type BaseTypeDescription string

const BaseTypeDescription_Native_Value BaseTypeDescription = "IdlNative"
const BaseTypeDescription_Mitch_Value BaseTypeDescription = "Mitch"

const (
	IDlBaseType_Native = BaseTypeDescription_Native_Value
	IDlBaseType_Mitch  = BaseTypeDescription_Mitch_Value
)

type IBaseTypeInformation interface {
	Name() BaseTypeDescription
	DefaultDecls() ([]IDefinitionDeclaration, error)
	CanScope(decl IDefinedType) bool
	CreateType(kind Kind, data interface{}) (IDefinedType, error)
}
