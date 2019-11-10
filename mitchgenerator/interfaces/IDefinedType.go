package interfaces

type IDefinedType interface {
	GetPackageName() (bool, string, string)
	GetSequenceCount() (bool, int)
	GetName() string
	Kind() Kind
	DefaultValue() string
	Predefined() bool
	GetStreamFunctionName() string
}
