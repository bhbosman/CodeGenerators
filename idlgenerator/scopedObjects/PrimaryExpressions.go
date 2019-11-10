package scopedObjects

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type PrimaryExpression struct {
	value     interface{}
	valuetype ScopingInterfaces.IPrimaryExpressionType
}

func NewPrimaryExpression(value interface{}, valuetype ScopingInterfaces.IPrimaryExpressionType) *PrimaryExpression {
	return &PrimaryExpression{
		value:     value,
		valuetype: valuetype,
	}
}

func (self *PrimaryExpression) Type() ScopingInterfaces.IPrimaryExpressionType {
	return self.valuetype
}

func (self *PrimaryExpression) Value() interface{} {
	return self.value
}
