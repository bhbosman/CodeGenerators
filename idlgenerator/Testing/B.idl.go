// Code generated by abc. DO NOT EDIT.

package Testing

//line B.idl:4
//Struct Decl: B
//Usage Count: 0
type B struct {
	C Common
}

//Constructors
func NewBDefaultPointer() (*B, error) {
	return &B{}, nil
}

func NewBDefaultValue() (B, error) {
	return B{}, nil
}

func NewBValue(
	c Common) (B, error) {
	return B{
		C: c,
	}, nil
}
