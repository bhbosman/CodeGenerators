package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

import "github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"

func Newconst_dcl() interfaces.IDefinitionDeclaration {
	return nil
}

func Newunion_dcl() interfaces.IDefinitionDeclaration {
	return nil
}

func Newstruct_forward_dcl() interfaces.IDefinitionDeclaration {
	return nil
}
