package interfaces

type IConstantValue interface {
	Value() interface{}
	ValueKind() Kind
	MaxLength() int
	//GetExportValue() string
}
