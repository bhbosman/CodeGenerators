package yacc

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

//go:generate goyacc -o completeIdl.go  -p "CompleteIdl"  completeIdl.y

type CompleteIdlLexer interface {
	Lex(lval *CompleteIdlSymType) int
	Error(s string)
	FindPrimitive(fileInformation ScopingInterfaces.IFileInformation, s string) (ScopingInterfaces.IDeclaredType, error)
	InfoAt(info string, params ...interface{})
	LastError() string
	MultiExpr(a int) (int, error)
	AddExpr(a int) (int, error)
	AddOperator(a, b int) (int, error)
	MinusOperator(a, b int) (int, error)
	DivideOperator(a, b int) (int, error)
	MultiplyOperator(a, b int) (int, error)
	NewStructType(id ScopingInterfaces.IFileInformation, identifier string, member ScopingInterfaces.IStructMember, forward bool) (ScopingInterfaces.IStructType, error)
	TransformString(string) (ScopingInterfaces.IPrimaryExpression, error)
	TransformInteger(int) (ScopingInterfaces.IPrimaryExpression, error)
	TransformValue(interface{}, ScopingInterfaces.IPrimaryExpressionType) (ScopingInterfaces.IPrimaryExpression, error)
	CreateTypePrefixDcl(fileInformation ScopingInterfaces.IFileInformation, scopedName, stringLiteral string) (ScopingInterfaces.ITypeSpec, error)
	CreateInterfaceDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, forward, abstract, local bool, body ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error)
	CreateModuleDcl(fileInformation ScopingInterfaces.IFileInformation, moduleName string, typeSpec ScopingInterfaces.ITypeSpec) (ScopingInterfaces.IIdlModuleDcl, error)
	AssignSpec(definition ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error)
	GetSpec() (ScopingInterfaces.ITypeSpec, error)
	NewDeclarator(fileInformation ScopingInterfaces.IFileInformation, identifier string) (ScopingInterfaces.IDeclarator, error)
	CreateInterfaceKind(fileInformation ScopingInterfaces.IFileInformation, local bool, abstract bool) (ScopingInterfaces.IInterfaceKind, error)
	NewMember(typeSpec ScopingInterfaces.IDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.IStructMember, error)
	NewIdlConstDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, value int) (ScopingInterfaces.IIdlConstDcl, error)
	NewTypeDeclarator(simpleTypeSpec ScopingInterfaces.IDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.ITypeDeclarator, error)
}

type IDefinitionContext interface {
	ParseExpression(string) bool
	ParseDefinition(string, bool)
	ParsePragma(string)
}
